package commands

import (
	"context"
	"os"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

func (Ops) BuildImage() error {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "docker")

	arch := os.Getenv("IMAGE_ARCH")
	if arch == "" {
		arch = "amd64"
	}

	tag := os.Getenv("IMAGE_TAG")
	if tag == "" {
		tag = "latest"
	}

	name := "wails-ubuntu-" + arch
	imageTag := "ghcr.io/kiloexabyte/" + name + ":" + tag
	dockerfile := "docker/Dockerfile.ubuntu-" + arch

	err := sh.Exec(ctx, "docker", "build",
		"-t", imageTag,
		"-f", dockerfile,
		".")
	if err != nil {
		return err
	}

	return sh.Exec(ctx, "docker", "images", imageTag,
		"--format", "Image Size: {{.Size}}")
}

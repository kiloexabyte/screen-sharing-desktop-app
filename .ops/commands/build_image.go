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

	tag := os.Getenv("IMAGE_TAG")
	if tag == "" {
		tag = "latest"
	}
	imageTag := "ghcr.io/kiloexabyte/wails-ubuntu:" + tag

	err := sh.Exec(ctx, "docker", "build",
		"-t", imageTag,
		"-f", "docker/Dockerfile.ubuntu",
		".")
	if err != nil {
		return err
	}

	return sh.Exec(ctx, "docker", "images", imageTag,
		"--format", "Image Size: {{.Size}}")
}

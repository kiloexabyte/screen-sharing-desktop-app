package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

func (Ops) UploadImage() error {
	ctx := context.Background()
	m := sys.Machine()

	user := os.Getenv("GHCR_USERNAME")
	token := os.Getenv("GHCR_TOKEN")

	// Login to ghcr.io
	loginStdin := command.NewWriter(ctx, m,
		"docker", "login", "ghcr.io",
		"-u", user,
		"--password-stdin",
	)

	if _, err := io.Copy(loginStdin, strings.NewReader(token)); err != nil {
		return fmt.Errorf("login stdin: %w", err)
	}

	if err := loginStdin.Close(); err != nil {
		return fmt.Errorf("docker login: %w", err)
	}

	// Push the image
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

	return command.Do(ctx, m, "docker", "push", imageTag)
}

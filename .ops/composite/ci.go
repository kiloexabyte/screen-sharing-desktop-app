package composite

import (
	"context"
	"fmt"

	"lesiw.io/command"
	"lesiw.io/command/sys"
	"lesiw.io/fs"
)

func (o Ops) Ci() error {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "npm")
	frontendCtx := fs.WithWorkDir(ctx, "frontend")

	fmt.Println("Building frontend...")
	if err := sh.Exec(frontendCtx, "npm", "install"); err != nil {
		return fmt.Errorf("npm install: %w", err)
	}
	if err := sh.Exec(frontendCtx, "npm", "run", "build"); err != nil {
		return fmt.Errorf("npm run build: %w", err)
	}

	fmt.Println("Running lint...")
	if err := o.Lint(); err != nil {
		return fmt.Errorf("lint: %w", err)
	}

	fmt.Println("Building all platforms...")
	if err := o.BuildAll(); err != nil {
		return fmt.Errorf("build_all: %w", err)
	}

	return nil
}

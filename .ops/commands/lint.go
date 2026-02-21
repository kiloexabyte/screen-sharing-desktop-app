package commands

import (
	"context"
	"fmt"

	"lesiw.io/command"
	"lesiw.io/command/sys"
	"lesiw.io/fs"
)

func (Ops) Lint() error {
	ctx := context.Background()

	// Lint Go code in root and .ops
	sh := command.Shell(sys.Machine(), "golangci-lint", "go")

	if err := sh.Exec(ctx, "golangci-lint", "run"); err != nil {
		return fmt.Errorf("golangci-lint (root): %w", err)
	}

	opsCtx := fs.WithWorkDir(ctx, ".ops")
	if err := sh.Exec(opsCtx, "golangci-lint", "run"); err != nil {
		return fmt.Errorf("golangci-lint (.ops): %w", err)
	}

	if err := sh.Exec(ctx, "go", "fmt", "./..."); err != nil {
		return fmt.Errorf("go fmt: %w", err)
	}

	// Lint frontend code
	sh = command.Shell(sys.Machine(), "npx", "npm")
	frontendCtx := fs.WithWorkDir(ctx, "frontend")

	if err := sh.Exec(frontendCtx, "npm", "install"); err != nil {
		return fmt.Errorf("npm install: %w", err)
	}

	glob := "./**/*.{js,jsx,mjs,cjs,ts,tsx,json,vue}"
	err := sh.Exec(frontendCtx, "npx", "prettier", "--check", glob)
	if err != nil {
		return fmt.Errorf("prettier: %w", err)
	}

	return nil
}

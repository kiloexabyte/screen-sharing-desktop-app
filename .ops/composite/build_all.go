package composite

import (
	"context"
	"fmt"
	"os"
	"strings"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

var defaultPlatforms = []string{
	"linux/amd64",
	"linux/386",
	"linux/arm64",
	"windows/amd64",
	"windows/386",
	"darwin/amd64",
	"darwin/arm64",
}

var crossCC = map[string]string{
	"linux/arm64":   "aarch64-linux-gnu-gcc",
	"windows/amd64": "x86_64-w64-mingw32-gcc",
}

func (Ops) BuildAll() error {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "wails")

	platforms := defaultPlatforms
	if env := os.Getenv("PLATFORMS"); env != "" {
		platforms = strings.Split(env, ",")
	}

	for _, platform := range platforms {
		fmt.Printf("Building for %s...\n", platform)
		buildCtx := ctx
		if cc, ok := crossCC[platform]; ok {
			buildCtx = command.WithEnv(ctx, map[string]string{
				"CC": cc,
			})
		}
		args := []string{
			"wails", "build",
			"-platform", platform,
		}
		if tags := os.Getenv("WAILS_TAGS"); tags != "" {
			args = append(args, "-tags", tags)
		}
		err := sh.Exec(buildCtx, args...)
		if err != nil {
			return err
		}
		fmt.Printf("Built %s successfully\n", platform)
	}
	return nil
}

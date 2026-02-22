package composite

import (
	"context"
	"fmt"
	"runtime"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

var platforms = []string{
	"linux/amd64",
	"linux/arm64",
	"windows/amd64",
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

	for _, platform := range platforms {
		fmt.Printf("Building for %s...\n", platform)
		buildCtx := ctx
		if cc, ok := crossCC[platform]; ok {
			buildCtx = command.WithEnv(ctx, map[string]string{
				"CC": cc,
			})
		}
		err := sh.Exec(
			buildCtx, "wails", "build",
			"-platform", platform,
		)
		if err != nil {
			isDarwin := len(platform) >= 6 &&
				platform[:6] == "darwin"
			if runtime.GOOS != "darwin" && isDarwin {
				fmt.Printf(
					"Skipping %s (requires macOS)\n",
					platform,
				)
				continue
			}
			return err
		}
		fmt.Printf("Built %s successfully\n", platform)
	}
	return nil
}

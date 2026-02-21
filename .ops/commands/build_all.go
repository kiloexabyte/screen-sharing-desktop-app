package commands

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

var platforms = []string{
	"linux/amd64",
	"windows/amd64",
	"darwin/amd64",
	"darwin/arm64",
}

func (Ops) BuildAll() {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "wails")

	for _, platform := range platforms {
		fmt.Printf("Building for %s...\n", platform)
		err := sh.Exec(ctx, "wails", "build", "-platform", platform)
		if err != nil {
			if runtime.GOOS != "darwin" && platform[:6] == "darwin" {
				fmt.Printf("Skipping %s (requires macOS)\n", platform)
				continue
			}
			log.Fatal(err)
		}
		fmt.Printf("Built %s successfully\n", platform)
	}
}

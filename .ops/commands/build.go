package commands

import (
	"context"
	"log"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

func (Ops) Build() {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "wails")

	if err := sh.Exec(ctx, "wails", "build", "-tags", "webkit2gtk_4.1"); err != nil {
		log.Fatal(err)
	}
}

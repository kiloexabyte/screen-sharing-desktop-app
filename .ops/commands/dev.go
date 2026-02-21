package commands

import (
	"context"
	"log"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

func (Ops) Dev() {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "wails")

	if err := sh.Exec(ctx, "wails", "dev"); err != nil {
		log.Fatal(err)
	}
}

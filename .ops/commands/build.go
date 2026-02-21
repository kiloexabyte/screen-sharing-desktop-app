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

	args := []string{"wails", "build", "-tags", "webkit2gtk_4.1"}
	if err := sh.Exec(ctx, args...); err != nil {
		log.Fatal(err)
	}
}

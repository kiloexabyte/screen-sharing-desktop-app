package commands

import (
	"context"

	"lesiw.io/command"
	"lesiw.io/command/sys"
)

func (Ops) Dev() error {
	ctx := context.Background()
	sh := command.Shell(sys.Machine(), "wails")
	return sh.Exec(ctx, "wails", "dev")
}

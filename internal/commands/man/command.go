package man

import (
	"fmt"

	"github.com/alecthomas/kong"
	mangokong "github.com/alecthomas/mango-kong"
	"github.com/muesli/roff"
)

type Man struct{}

func (m Man) BeforeApply(ctx *kong.Context) error {
	// Set the correct man pages description without color escape sequences.
	ctx.Model.Help = "A tool to apply commitlint to your commits ✨"
	man := mangokong.NewManPage(1, ctx.Model)
	man = man.WithSection("Copyright", "(c) 2025-2026 Jurien Hamaker.\n"+
		"Released under MIT license.")
	_, _ = fmt.Fprint(ctx.Stdout, man.Build(roff.NewDocument()))
	ctx.Exit(0)
	return nil
}

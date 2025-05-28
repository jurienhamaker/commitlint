package commitlint

import (
	"github.com/alecthomas/kong"

	"github.com/jurienhamaker/commitlint/internal/commands/install"
	"github.com/jurienhamaker/commitlint/internal/commands/lint"
	"github.com/jurienhamaker/commitlint/internal/commands/man"
)

type Commitlint struct {
	// Version is a flag that can be used to display the version number.
	Version kong.VersionFlag `short:"v" help:"Print the version number"`

	// Man is a hidden command that generates Gum man pages.
	Man man.Man `cmd:"" hidden:"" help:"Generate man pages"`

	// Lint is the main command, it will check the given message for
	// the applied rules
	//
	// Apply linting on a message:
	//
	// $ commitlint lint "fix: something"
	// $ commitlint "fix: something"
	// $ echo "fix: something" | commitlint
	// $ echo "fix: something" | commitlint lint
	//
	Lint lint.Lint `cmd:"" default:"withargs" help:"Test the given message"`

	// Install is a command to install commitlint on your git repository.
	// This will add a pre-commit hook to your repository that will run
	// before every commit and apply commitlint to the commit message
	//
	// Install commitlint in your repository:
	//
	// $ commitlint install
	//
	Install install.Install `cmd:"" help:"Install commitlint in your repository"`
}

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"github.com/muesli/termenv"

	"github.com/jurienhamaker/commitlint/internal/commitlint"
	"github.com/jurienhamaker/commitlint/internal/exit"
)

const shaLen = 7

var (
	// Version contains the application version number. It's set via ldflags
	// when building.
	Version = ""

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via ldflags when building.
	CommitSHA = ""
)

func main() {
	lipgloss.SetColorProfile(termenv.NewOutput(os.Stderr).Profile)

	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
			Version = info.Main.Version
		} else {
			Version = "unknown (built from source)"
		}
	}

	version := fmt.Sprintf("commitlint version %s", Version)
	if len(CommitSHA) >= shaLen {
		version += " (" + CommitSHA[:shaLen] + ")"
	}

	cli := &commitlint.Commitlint{}
	ctx := kong.Parse(
		cli,
		kong.Description("A tool to apply commitlint to your commits ✨"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             true,
			Summary:             false,
			NoExpandSubcommands: true,
		}),
		kong.Vars{
			"version":                 version,
			"versionNumber":           Version,
			"defaultHeight":           "0",
			"defaultWidth":            "0",
			"defaultAlign":            "left",
			"defaultBorder":           "none",
			"defaultBorderForeground": "",
			"defaultBorderBackground": "",
			"defaultBackground":       "",
			"defaultForeground":       "",
			"defaultMargin":           "0 0",
			"defaultPadding":          "0 0",
			"defaultUnderline":        "false",
			"defaultBold":             "false",
			"defaultFaint":            "false",
			"defaultItalic":           "false",
			"defaultStrikethrough":    "false",
		},
	)

	if err := ctx.Run(); err != nil {
		var ex exit.ErrExit
		if errors.As(err, &ex) {
			os.Exit(int(ex))
		}

		if errors.Is(err, tea.ErrProgramKilled) {
			log.Error(os.Stderr, "timed out")
			os.Exit(exit.StatusTimeout)
		}

		if errors.Is(err, tea.ErrInterrupted) {
			os.Exit(exit.StatusAborted)
		}

		log.Error(os.Stderr, err)
		os.Exit(1)
	}
}

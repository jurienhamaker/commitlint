package lint

import (
	"fmt"
	"strings"

	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/validation"
)

type ParseResult map[validation.ValidationState]int

func parseResult(message string, config *config.Config, result validation.ValidationsResult) (total int, parseResult ParseResult) {
	fmt.Println("\n")

	success := result[validation.ValidationStateSuccess]
	warning := result[validation.ValidationStateWarning]
	error := result[validation.ValidationStateError]
	parseResult = ParseResult{
		validation.ValidationStateSuccess: len(success),
		validation.ValidationStateWarning: len(warning),
		validation.ValidationStateError:   len(error),
	}
	total = len(success) + len(warning) + len(error)

	if total == 0 {
		fmt.Printf("No rules have been checked")
	}

	emoji := "💬"
	if !config.UseEmoji {
		emoji = "✧"
	}

	fmt.Printf("  %s\n", styles.SupportiveLilacTextStyle(fmt.Sprintf("%s %s", emoji, strings.ReplaceAll(message, "\\n", "\n\t "))))
	printStateMessages(success, config, validation.ValidationStateSuccess)
	printStateMessages(warning, config, validation.ValidationStateWarning)
	printStateMessages(error, config, validation.ValidationStateError)

	fmt.Println(" ")

	return
}

func printStateMessages(messages validation.ValidationResult, config *config.Config, state validation.ValidationState) {
	for _, result := range messages {
		textStyle := styles.ValidationStateStyle[state]
		emoji := styles.ValidationStateEmoji[state]
		if !config.UseEmoji {
			emoji = styles.ValidationStateUnicode[state]
		}

		fmt.Printf("  %s %s\n", textStyle(fmt.Sprintf("%s %s", emoji, result.Message)), styles.GrayishTextStyle(fmt.Sprintf("[%s]", result.Rule)))
	}
}

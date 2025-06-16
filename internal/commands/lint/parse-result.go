package lint

import (
	"fmt"
	"strings"

	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/validation"
)

type ParseResult map[validation.ValidationState]int

func parseResult(message string, result validation.ValidationsResult) (total int, parseResult ParseResult) {
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

	fmt.Printf("  💬 %s\n", styles.SupportiveLilacTextStyle(strings.ReplaceAll(message, "\\n", "\n\t ")))
	printStateMessages(success, validation.ValidationStateSuccess)
	printStateMessages(warning, validation.ValidationStateWarning)
	printStateMessages(error, validation.ValidationStateError)

	fmt.Println(" ")

	return
}

func printStateMessages(messages validation.ValidationResult, state validation.ValidationState) {
	for _, result := range messages {
		textStyle := styles.ValidationStateStyle[state]
		emoji := styles.ValidationStateEmoji[state]
		fmt.Printf("  %s %s %s\n", emoji, textStyle(result.Message), styles.GrayishTextStyle(fmt.Sprintf("[%s]", result.Rule)))
	}
}

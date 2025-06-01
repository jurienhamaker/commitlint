package lint

import (
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/validation"
)

type ParseResult map[validation.ValidationState]int

func parseResult(result validation.ValidationsResult, message string) (total int, parseResult ParseResult) {
	fmt.Println("\nCommitlint result:")

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

	printStateMessages(success, validation.ValidationStateSuccess)
	printStateMessages(warning, validation.ValidationStateWarning)
	printStateMessages(error, validation.ValidationStateError)

	fmt.Println(" ")

	return
}

func printStateMessages(messages []string, state validation.ValidationState) {
	for _, str := range messages {
		textStyle := styles.ValidationStateStyle[state]
		emoji := styles.ValidationStateEmoji[state]
		fmt.Printf("  %s %s\n", emoji, textStyle(str))
	}
}

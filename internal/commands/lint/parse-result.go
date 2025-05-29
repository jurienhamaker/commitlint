package lint

import (
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/validation"
)

type ParseResult map[validation.ValidationState]int

func parseResult(result validation.ValidationsResult, message string) (parseResult ParseResult) {
	fmt.Printf("\nCommitlint result for \"%s\":\n", styles.BoldTextStyle(message))

	parseResult = make(ParseResult)

	for pluginName, validations := range result {
		if len(validations) == 0 {
			continue
		}

		fmt.Printf("  • %s:\n", styles.BoldTextStyle(pluginName))
		for str, state := range validations {
			parseResult[state]++

			textStyle := styles.ValidationStateStyle[state]
			emoji := styles.ValidationStateEmoji[state]
			fmt.Printf("    %s %s\n", emoji, textStyle(str))
		}
	}

	fmt.Println(" ")

	return
}

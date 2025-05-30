package lint

import (
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/styles"
	"github.com/jurienhamaker/commitlint/validation"
)

type ParseResult map[validation.ValidationState]int

func parseResult(result validation.ValidationsResult, message string) (total int, parseResult ParseResult) {
	fmt.Printf("\nCommitlint result for \"%s\":\n", styles.BoldTextStyle(message))

	parseResult = make(ParseResult)

	if len(result) == 0 {
		fmt.Printf("No rules have been checked")
	}

	for _, result := range result {
		for str, state := range result {
			total++
			parseResult[state]++

			textStyle := styles.ValidationStateStyle[state]
			emoji := styles.ValidationStateEmoji[state]
			fmt.Printf("  %s %s\n", emoji, textStyle(str))
		}
	}

	fmt.Println(" ")

	return
}

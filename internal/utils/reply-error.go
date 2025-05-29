package utils

import (
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/styles"
)

func ReplyError(message string) {
	fmt.Println(
		styles.ErrorStyle(
			fmt.Sprintf("\nError: %s\n", message),
		),
	)
}

func ReplyWarning(message string) {
	fmt.Println(
		styles.WarningTextStyle(
			fmt.Sprintf("\nWarning: %s\n", message),
		),
	)
}

func ReplySuccess(message string) {
	fmt.Println(
		styles.SuccessTextStyle(
			fmt.Sprintf("\nSuccess: %s\n", message),
		),
	)
}

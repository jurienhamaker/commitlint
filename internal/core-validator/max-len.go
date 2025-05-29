package corevalidator

import (
	"fmt"

	"github.com/jurienhamaker/commitlint/validation"
)

func maxLenValidator(message string, amount int, level validation.ValidationState) (rule string, state validation.ValidationState) {
	state = validation.ValidationStateSuccess
	rule = fmt.Sprintf("Message must be less than %d characters", amount)

	if len(message) > amount {
		state = level
	}

	return
}

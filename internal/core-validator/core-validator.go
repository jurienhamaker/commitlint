package corevalidator

import (
	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

func CoreValidator(commit *parser.ConventionalCommit, config validation.ValidatorConfig) (result validation.ValidationResult, err error) {
	result = make(validation.ValidationResult)

	bodyCaseConf, ok := config["body-case"]
	if ok {
		message, state, resultErr := bodyCase(commit, bodyCaseConf)
		if resultErr != nil {
			err = resultErr
			return
		}

		result[message] = state
	}

	bodyFullStopConf, ok := config["body-full-stop"]
	if ok {
		message, state, resultErr := bodyFullStop(commit, bodyFullStopConf)
		if resultErr != nil {
			err = resultErr
			return
		}

		result[message] = state
	}

	return
}

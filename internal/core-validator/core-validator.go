package corevalidator

import (
	"maps"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"

	"github.com/jurienhamaker/commitlint/internal/core-validator/validators"
)

func CoreValidator(commit *parser.ConventionalCommit, config validation.ValidatorConfig) (result validation.ValidationResult, err error) {
	result = make(validation.ValidationResult)

	validators := map[string]ValidatorFn{
		"body-case":              validators.BodyCase,
		"body-empty":             validators.BodyEmpty,
		"body-full-stop":         validators.BodyFullStop,
		"body-leading-blank":     validators.BodyLeadingBlank,
		"body-max-length":        validators.BodyMaxLength,
		"body-max-line-length":   validators.BodyMaxLineLength,
		"body-min-length":        validators.BodyMinLength,
		"footer-empty":           validators.FooterEmpty,
		"footer-leading-blank":   validators.FooterLeadingBlank,
		"footer-max-length":      validators.FooterMaxLength,
		"footer-max-line-length": validators.FooterMaxLineLength,
		"footer-min-length":      validators.FooterMinLength,
	}

	for name, fn := range validators {
		validatorResult, validatorErr := checkValidator(
			commit,
			name,
			config,
			fn,
		)
		if validatorErr != nil {
			err = validatorErr
			return
		}
		maps.Copy(result, validatorResult)
	}

	return
}

func checkValidator(commit *parser.ConventionalCommit, name string, config validation.ValidatorConfig, fn ValidatorFn) (result validation.ValidationResult, err error) {
	result = make(validation.ValidationResult)
	conf, ok := config[name]
	if ok {
		message, state, resultErr := fn(commit, conf)
		if resultErr != nil {
			err = resultErr
			return
		}

		result[message] = state
	}

	return
}

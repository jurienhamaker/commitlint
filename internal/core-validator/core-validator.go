package corevalidator

import (
	"maps"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"

	"github.com/jurienhamaker/commitlint/internal/core-validator/rules"
)

func CoreValidator(commit *parser.ConventionalCommit, config validation.ValidatorConfig) (result validation.ValidationResult, err error) {
	result = make(validation.ValidationResult)

	rules := map[string]ValidatorFn{
		"body-case":              rules.BodyCase,
		"body-empty":             rules.BodyEmpty,
		"body-full-stop":         rules.BodyFullStop,
		"body-leading-blank":     rules.BodyLeadingBlank,
		"body-max-length":        rules.BodyMaxLength,
		"body-max-line-length":   rules.BodyMaxLineLength,
		"body-min-length":        rules.BodyMinLength,
		"footer-empty":           rules.FooterEmpty,
		"footer-leading-blank":   rules.FooterLeadingBlank,
		"footer-max-length":      rules.FooterMaxLength,
		"footer-max-line-length": rules.FooterMaxLineLength,
		"footer-min-length":      rules.FooterMinLength,
		"header-case":            rules.HeaderCase,
		"header-full-stop":       rules.HeaderFullStop,
		"header-max-length":      rules.HeaderMaxLength,
		"header-min-length":      rules.HeaderMinLength,
		"header-trim":            rules.HeaderTrim,
		"references":             rules.References,
		"scope-case":             rules.ScopeCase,
		"scope-empty":            rules.ScopeEmpty,
		"scope-enum":             rules.ScopeEnum,
		"scope-max-length":       rules.ScopeMaxLength,
		"scope-min-length":       rules.ScopeMinLength,
		"signed-off-by":          rules.SignedOffBy,
		"subject-case":           rules.SubjectCase,
		"subject-empty":          rules.SubjectEmpty,
		"subject-full-stop":      rules.SubjectFullStop,
		"subject-max-length":     rules.SubjectMaxLength,
		"subject-min-length":     rules.SubjectMinLength,
		"type-case":              rules.TypeCase,
		"type-empty":             rules.TypeEmpty,
		"type-enum":              rules.TypeEnum,
		"type-max-length":        rules.TypeMaxLength,
		"type-min-length":        rules.TypeMinLength,
	}

	for name, fn := range rules {
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

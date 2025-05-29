package corevalidator

import (
	"errors"
	"maps"
	"reflect"

	"github.com/charmbracelet/log"
	"github.com/jurienhamaker/commitlint/validation"
)

var BASE_CONFIG = map[string]map[string]any{
	"max-len": {
		"enabled": true,
		"value":   144,
		"level":   validation.ValidationStateError,
	},
	"min-len": {
		"enabled": true,
		"value":   5,
		"level":   validation.ValidationStateError,
	},
}

func getRuleConfig(ruleName string, config map[string]any) map[string]any {
	v := BASE_CONFIG[ruleName]
	ruleConfig := config[ruleName]

	if reflect.ValueOf(ruleConfig).Kind() == reflect.Bool {
		if !ruleConfig.(bool) {
			maps.Copy(v, map[string]any{
				"enabled": false,
			})
		}

		return v
	}

	maps.Copy(v, ruleConfig.(map[string]any))
	return v
}

func CoreValidator(message string, config validation.ValidatorConfig) (result validation.ValidationResult, err error) {
	log.Debugf("Core validator running for: %s", message)

	result = make(validation.ValidationResult)

	err = maxLen(message, result, config["max-len"])
	if err != nil {
		return
	}

	err = minLen(message, result, config["min-len"])
	if err != nil {
		return
	}

	return
}

func maxLen(message string, result validation.ValidationResult, config validation.ValidationRuleConfig) error {
	if !config.Enabled {
		return nil
	}

	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		return errors.New("max-len rule value must be an integer")
	}

	maxLen, state := maxLenValidator(message, config.Value.(int), config.Level)
	if maxLen != "" {
		result[maxLen] = state
	}

	return nil
}

func minLen(message string, result validation.ValidationResult, config validation.ValidationRuleConfig) error {
	if !config.Enabled {
		return nil
	}

	if reflect.ValueOf(config.Value).Kind() != reflect.Int {
		return errors.New("min-len rule value must be an integer")
	}

	minLen, state := minLenValidator(message, config.Value.(int), config.Level)
	if minLen != "" {
		result[minLen] = state
	}

	return nil
}

package corevalidator

import (
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

func CoreValidator(message string, config map[string]any) (result validation.ValidationResult, err error) {
	log.Debugf("Core validator running for: %s", message)

	result = make(validation.ValidationResult)

	maxLenConfig := getRuleConfig("max-len", config)
	maxLen(message, result, maxLenConfig)

	minLenConfig := getRuleConfig("min-len", config)
	minLen(message, result, minLenConfig)

	return
}

func maxLen(message string, result validation.ValidationResult, config map[string]any) {
	if !config["enabled"].(bool) {
		return
	}

	level := config["level"]
	if reflect.ValueOf(level).Kind() == reflect.String {
		level = validation.ValidationStateInt[level.(string)]
	}

	maxLen, state := maxLenValidator(message, config["value"].(int), level.(validation.ValidationState))
	if maxLen != "" {
		result[maxLen] = state
	}
}

func minLen(message string, result validation.ValidationResult, config map[string]any) {
	if !config["enabled"].(bool) {
		return
	}

	level := config["level"]
	if reflect.ValueOf(level).Kind() == reflect.String {
		level = validation.ValidationStateInt[level.(string)]
	}

	maxLen, state := minLenValidator(message, config["value"].(int), level.(validation.ValidationState))
	if maxLen != "" {
		result[maxLen] = state
	}
}

package config

import (
	"errors"
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/constants"
	"github.com/jurienhamaker/commitlint/validation"
	"github.com/spf13/viper"
)

type (
	Rules  map[string]validation.ValidatorConfig
	Config struct {
		Enabled bool
		Rules   Rules
	}
)

var c Config

func init() {
	viper.SetConfigName(constants.CONFIG_NAME)
	viper.SetConfigType(constants.CONFIG_TYPE)
	viper.AddConfigPath(constants.CONFIG_PATH)
}

func parseRuleConfig(ruleConfig []any) (c validation.ValidationRuleConfig, err error) {
	c = validation.ValidationRuleConfig{
		Enabled: true,
		Level:   validation.ValidationStateError,
		Value:   nil,
	}

	if len(ruleConfig) == 0 {
		return
	}

	enabled, ok := ruleConfig[0].(string)
	if ok && enabled == "never" {
		c.Enabled = false
	}

	levelInt, ok := ruleConfig[0].(int)
	if !ok || levelInt < 0 || levelInt > 2 {
		err = errors.New("first option of rule must be a number between 0 and 2")
		return
	}

	level, _ := ruleConfig[0].(validation.ValidationState)
	c.Level = level

	enabled, ok = ruleConfig[1].(string)
	if !ok || (enabled != "never" && enabled != "always") {
		err = fmt.Errorf("second option of rule must be 'never' or 'always', found \"%s\"", enabled)
		return
	}

	if enabled == "never" {
		c.Enabled = false
	}

	if len(ruleConfig) > 2 {
		c.Value = ruleConfig[2]
	}

	return
}

func parseConfig() (Config, error) {
	c = Config{}
	c.Enabled = viper.GetBool("enabled")

	c.Rules = make(Rules)
	rules := viper.GetStringMap("rules")

	for pluginName := range rules {
		ruleConfig := viper.GetStringMap(fmt.Sprintf("rules.%s", pluginName))
		c.Rules[pluginName] = make(validation.ValidatorConfig)

		for rule, ruleConfig := range ruleConfig {
			parsed, err := parseRuleConfig(ruleConfig.([]any))
			if err != nil {
				panic(fmt.Errorf("Config [rules.%s.%s] %s", pluginName, rule, err.Error()))
			}

			c.Rules[pluginName][rule] = parsed
		}
	}

	return c, nil
}

func Load() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	_, err = parseConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return c
}

package config

import (
	"errors"
	"fmt"

	"github.com/jurienhamaker/commitlint/internal/constants"
	"github.com/jurienhamaker/commitlint/validation"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Enabled  bool
		UseEmoji bool
		Rules    validation.ValidatorConfig
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
		Always: true,
		Level:  validation.ValidationStateError,
		Value:  nil,
	}

	if len(ruleConfig) == 0 {
		return
	}

	level, ok := ruleConfig[0].(int)
	if !ok || level < 0 || level > 2 {
		err = errors.New("first option of rule must be a number between 0 and 2")
		return
	}

	c.Level = validation.ValidationStateMapping[level]

	if len(ruleConfig) == 1 {
		c.Always = true
		return
	}

	always, ok := ruleConfig[1].(string)
	if !ok || (always != "never" && always != "always") {
		err = fmt.Errorf("second option of rule must be 'never' or 'always', found \"%s\"", always)
		return
	}

	if always == "never" {
		c.Always = false
	}

	if len(ruleConfig) > 2 {
		c.Value = ruleConfig[2]
	}

	return
}

func parseConfig() Config {
	c = Config{}
	c.Enabled = viper.GetBool("enabled")
	c.UseEmoji = viper.GetBool("use-emoji")

	c.Rules = make(validation.ValidatorConfig)
	rules := viper.GetStringMap("rules")

	for rule, ruleConfig := range rules {
		parsed, err := parseRuleConfig(ruleConfig.([]any))
		if err != nil {
			panic(fmt.Errorf("Config [rules.%s] %s", rule, err.Error()))
		}

		c.Rules[rule] = parsed
	}

	return c
}

func Load() error {
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read config: %s", err)
	}

	parseConfig()

	return nil
}

func GetConfig() *Config {
	return &c
}

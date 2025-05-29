package config

import (
	"github.com/jurienhamaker/commitlint/internal/constants"
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

var c *Config

func init() {
	c = new(Config)
	c.Viper = viper.GetViper()

	c.Viper.SetConfigName(constants.CONFIG_NAME)
	c.Viper.SetConfigType(constants.CONFIG_TYPE)
	c.Viper.AddConfigPath(constants.CONFIG_PATH)
}

func Load() (*Config, error) {
	err := c.Viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetConfig() *Config {
	return c
}

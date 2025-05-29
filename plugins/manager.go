// Package plugin serves as the bridge between the main application and plugins.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugins

import (
	"github.com/charmbracelet/log"
	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/validation"
)

type PluginManager struct {
	plugins map[string]validation.Validator
}

func newPluginManager() *PluginManager {
	pm := &PluginManager{}
	pm.plugins = make(map[string]validation.Validator)
	return pm
}

func (pm *PluginManager) RegisterPlugin(pluginName string, validator validation.Validator) {
	pm.plugins[pluginName] = validator
}

func (pm *PluginManager) RunPluginValidators(message string) (results validation.ValidationsResult, err error) {
	log.Debugf("Running plugins: %s", message)
	results = make(validation.ValidationsResult)

	c := config.GetConfig()

	for pluginName, validator := range pm.plugins {
		pluginConfig := c.Rules[pluginName]
		result, valErr := validator(message, pluginConfig)
		if valErr != nil {
			err = valErr
			break
		}

		results[pluginName] = result
	}

	return
}

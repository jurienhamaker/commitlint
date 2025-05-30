// Package plugin serves as the bridge between the main application and plugins.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugins

import (
	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/parser"
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

func (pm *PluginManager) RunPluginValidators(commit *parser.ConventionalCommit) (results validation.ValidationsResult, err error) {
	results = validation.ValidationsResult{}

	c := config.GetConfig()

	for _, validator := range pm.plugins {
		result, valErr := validator(commit, c.Rules)
		if valErr != nil {
			err = valErr
			break
		}

		results = append(results, result)
	}

	return
}

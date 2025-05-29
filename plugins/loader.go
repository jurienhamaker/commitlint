// Plugin loading code.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugins

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

// LoadPlugins loads plugins from the directory with the given path, looking for
// all .so files in there. It creates a new PluginManager and registers the
// plugins with it.
func LoadPlugins(path string) (*PluginManager, error) {
	pm := newPluginManager()

	c, err := os.ReadDir(path)
	if os.IsNotExist(err) {
		// just return manager because there is no plugins folder so we won't load anything
		return pm, nil
	}

	if err != nil {
		return nil, err
	}

	for _, entry := range c {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".so") {
			fullpath := filepath.Join(path, entry.Name())
			fmt.Println("found plugin file", fullpath)

			p, err := plugin.Open(fullpath)
			if err != nil {
				return nil, err
			}

			ifunc, err := p.Lookup("InitPlugin")
			if err != nil {
				return nil, err
			}

			initFunc := ifunc.(func(*PluginManager) error)
			if err := initFunc(pm); err != nil {
				return nil, err
			}
		}
	}
	return pm, nil
}

package constants

import "fmt"

const (
	DEFAULT_CONFIG = `enabled: true

rules: []`
	CONFIG_TYPE = "yaml"
	CONFIG_NAME = "commitlint"
	CONFIG_PATH = ".commitlint"
)

var (
	CONFIG_FILE_NAME = fmt.Sprintf("%s.%s", CONFIG_NAME, CONFIG_TYPE)
	CONFIG_FILE_PATH = fmt.Sprintf("%s/%s", CONFIG_PATH, CONFIG_FILE_NAME)
)

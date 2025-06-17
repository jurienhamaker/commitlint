package constants

import "fmt"

const (
	DEFAULT_CONFIG = `enabled: true
use-emoji: true

rules:
  header-max-length: [2, "always", 100]
  header-trim: [2]
  type-case: [2, "always", "lowercase"]
  type-empty: [2, "never"]
  type-enum: [2, "always", ["build", "chore", "ci", "docs", "feat", "fix", "perf", "refactor", "revert", "style", "test"]]
  subject-case: [2, "never", ["pascal-case", "upper-case"]]
  subject-empty: [2, "never"]
  subject-full-stop: [2, "never", "."]
  body-leading-blank: [1, "always"]
  body-max-line-length: [2, "always", 100]
  footer-leading-blank: [1, "always"]
  footer-max-line-length: [2, "always", 100]`
	CONFIG_TYPE = "yaml"
	CONFIG_NAME = "commitlint"
	CONFIG_PATH = ".commitlint"
)

var (
	CONFIG_FILE_NAME = fmt.Sprintf("%s.%s", CONFIG_NAME, CONFIG_TYPE)
	CONFIG_FILE_PATH = fmt.Sprintf("%s/%s", CONFIG_PATH, CONFIG_FILE_NAME)
)

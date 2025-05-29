package constants

import "fmt"

const (
	GIT_PATH = ".git"
)

var (
	HOOKS_PATH      = fmt.Sprintf("%s/hooks", GIT_PATH)
	COMMIT_MSG_PATH = fmt.Sprintf("%s/commit-msg", HOOKS_PATH)
)

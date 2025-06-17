package constants

import "fmt"

const (
	GIT_PATH        = ".git"
	HOOKS_DIRECTORY = "hooks"
	COMMIT_MSG_FILE = "commit-msg"
)

var (
	HOOKS_PATH_LOCAL      = fmt.Sprintf("%s/%s", GIT_PATH, HOOKS_DIRECTORY)
	COMMIT_MSG_PATH_LOCAL = fmt.Sprintf("%s/%s", HOOKS_PATH_LOCAL, COMMIT_MSG_FILE)

	COMMIT_MSG_PATH = fmt.Sprintf("%s/%s", HOOKS_DIRECTORY, COMMIT_MSG_FILE)
)

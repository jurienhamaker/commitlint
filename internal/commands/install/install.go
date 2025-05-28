package install

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const hookCommand = "cat $1 | commitlint"

var hookContent = fmt.Sprintf(`#!/bin/sh
%s`, hookCommand)
var hookContentMessage = fmt.Sprintf("Please add the following line to the hook:\n\n%s\n", hookCommand)

func ensureHooksDirectory() error {
	gitStat, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return errors.New("no .git folder found in current directory")
	}

	if !gitStat.IsDir() {
		return errors.New(".git is not a folder")
	}

	hooksStat, err := os.Stat(".git/hooks")
	created := false
	if os.IsNotExist(err) {
		err = os.Mkdir(".git/hooks", 0755)
		created = true
		if err != nil {
			return errors.New("could not create .git/hooks folder")
		}
	}

	if !created && !hooksStat.IsDir() {
		return errors.New(".git/hooks is not a folder")
	}

	return nil
}

func hookExists() error {
	_, err := os.Stat(".git/hooks/commit-msg")

	if !os.IsNotExist(err) {
		return fmt.Errorf(".git/hooks/commit-msg already exists.\n%s", hookContentMessage)
	}

	return nil
}

func install(sub chan resultMsg) {
	err := ensureHooksDirectory()
	if err != nil {
		sub <- resultMsg{Error: err}
		return
	}

	err = hookExists()
	if err != nil {
		sub <- resultMsg{Error: err}
		return
	}

	command := []byte(hookContent)
	err = os.WriteFile(".git/hooks/commit-msg", command, 0744)
	if err != nil {
		sub <- resultMsg{Error: fmt.Errorf("couldn't create .git/hooks/commit-msg.\n%s", hookContentMessage)}
	}

	time.Sleep(time.Second * 1)
	sub <- resultMsg{Installed: true}
}

package install

import (
	"fmt"
	"os"
	"time"

	"github.com/jurienhamaker/commitlint/internal/constants"
	"github.com/jurienhamaker/commitlint/internal/spinner"
)

const hookCommand = "cat $1 | commitlint"

var hookContent = fmt.Sprintf(`#!/bin/sh
%s`, hookCommand)
var hookContentMessage = fmt.Sprintf("Please add the following line to the hook:\n\n%s", hookCommand)

func ensureHooksDirectory() error {
	gitStat, err := os.Stat(constants.GIT_PATH)
	if os.IsNotExist(err) {
		return fmt.Errorf("no %s folder found in current directory", constants.GIT_PATH)
	}

	if !gitStat.IsDir() {
		return fmt.Errorf("%s is not a folder", constants.GIT_PATH)
	}

	hooksStat, err := os.Stat(constants.HOOKS_PATH)
	created := false
	if os.IsNotExist(err) {
		err = os.Mkdir(constants.HOOKS_PATH, 0755)
		created = true
		if err != nil {
			return fmt.Errorf("could not create %s folder", constants.HOOKS_PATH)
		}
	}

	if !created && !hooksStat.IsDir() {
		return fmt.Errorf("%s is not a folder", constants.HOOKS_PATH)
	}

	return nil
}

func hookExists() error {
	_, err := os.Stat(constants.COMMIT_MSG_PATH)

	if !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists.\n%s", constants.COMMIT_MSG_PATH, hookContentMessage)
	}

	return nil
}

func checkConfigFile() error {
	configStat, err := os.Stat(constants.CONFIG_PATH)
	created := false
	if os.IsNotExist(err) {
		err = os.Mkdir(constants.CONFIG_PATH, 0755)
		created = true
		if err != nil {
			return fmt.Errorf("could not create %s folder", constants.CONFIG_PATH)
		}
	}

	if !created && !configStat.IsDir() {
		return fmt.Errorf("%s is not a folder", constants.CONFIG_PATH)
	}

	_, err = os.Stat(constants.CONFIG_FILE_PATH)
	if os.IsNotExist(err) {
		//  we will create the default config
		command := []byte(constants.DEFAULT_CONFIG)
		err = os.WriteFile(constants.CONFIG_FILE_PATH, command, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func install(sub chan spinner.SpinnerResultMsg[bool]) {
	err := checkConfigFile()
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	err = ensureHooksDirectory()
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	err = hookExists()
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	command := []byte(hookContent)
	err = os.WriteFile(constants.COMMIT_MSG_PATH, command, 0744)
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: fmt.Errorf("couldn't create %s.\n%s", constants.COMMIT_MSG_PATH, hookContentMessage)}
	}

	time.Sleep(time.Second * 1)
	sub <- spinner.SpinnerResultMsg[bool]{Result: true}
}

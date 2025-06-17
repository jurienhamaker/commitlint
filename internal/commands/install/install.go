package install

import (
	"fmt"
	"os"
	"time"

	"github.com/jurienhamaker/commitlint/config"
	"github.com/jurienhamaker/commitlint/internal/constants"
	"github.com/jurienhamaker/commitlint/internal/spinner"
)

const hookCommand = "cat $1 | commitlint"

var hookContent = fmt.Sprintf(`#!/bin/sh
%s`, hookCommand)
var hookContentMessage = fmt.Sprintf("Please add the following line to the hook:\n\n%s", hookCommand)

func getHooksPath(global bool) (string, error) {
	hooksPath := constants.HOOKS_PATH_LOCAL

	if global {
		globalPath, err := config.GetGlobalPath()
		if err != nil {
			return "", fmt.Errorf("could not load global path: %s", err)
		}

		hooksPath = fmt.Sprintf("%s/%s", globalPath, constants.HOOKS_DIRECTORY)
	}

	return hooksPath, nil
}

func ensureHooksDirectory(global bool) error {
	if !global {
		gitStat, err := os.Stat(constants.GIT_PATH)
		if os.IsNotExist(err) {
			return fmt.Errorf("no %s folder found in current directory", constants.GIT_PATH)
		}

		if !gitStat.IsDir() {
			return fmt.Errorf("%s is not a folder", constants.GIT_PATH)
		}
	}

	hooksPath, err := getHooksPath(global)
	if err != nil {
		return err
	}

	hooksStat, err := os.Stat(hooksPath)
	created := false
	if os.IsNotExist(err) {
		err = os.Mkdir(hooksPath, 0o750)
		created = true
		if err != nil {
			return fmt.Errorf("could not create %s folder", hooksPath)
		}
	}

	if !created && !hooksStat.IsDir() {
		return fmt.Errorf("%s is not a folder", hooksPath)
	}

	return nil
}

func getHookFilePath(global bool) (string, error) {
	hooksPath, err := getHooksPath(global)
	if err != nil {
		return "", err
	}

	hookFilePath := fmt.Sprintf("%s/%s", hooksPath, constants.COMMIT_MSG_FILE)
	return hookFilePath, nil
}

func hookExists(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists.\n%s", constants.COMMIT_MSG_PATH_LOCAL, hookContentMessage)
	}

	return nil
}

func checkConfigFile(global bool) error {
	configPath, err := config.GetPath(global)
	if err != nil {
		return fmt.Errorf("could not get path: %s", err)
	}

	configStat, err := os.Stat(configPath)
	created := false
	if os.IsNotExist(err) {
		err = os.Mkdir(configPath, 0o750)
		created = true
		if err != nil {
			return fmt.Errorf("could not create %s folder", constants.CONFIG_PATH)
		}
	}

	if !created && !configStat.IsDir() {
		return fmt.Errorf("%s is not a folder", constants.CONFIG_PATH)
	}

	configFilePath, err := config.GetFilePath(global)
	if err != nil {
		return fmt.Errorf("could not get filepath: %s", err)
	}

	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		//  we will create the default config
		command := []byte(constants.DEFAULT_CONFIG)
		err = os.WriteFile(configFilePath, command, 0o600)
		if err != nil {
			return fmt.Errorf("could not create config file: %s", err)
		}
	}

	return nil
}

func install(sub chan spinner.SpinnerResultMsg[bool], global bool) {
	err := checkConfigFile(global)
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	err = ensureHooksDirectory(global)
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	hookFilePath, err := getHookFilePath(global)
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	err = hookExists(hookFilePath)
	if err != nil && global {
		time.Sleep(time.Second * 1)
		sub <- spinner.SpinnerResultMsg[bool]{Result: true}
	}

	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: err}
		return
	}

	command := []byte(hookContent)
	err = os.WriteFile(hookFilePath, command, 0o700)
	if err != nil {
		sub <- spinner.SpinnerResultMsg[bool]{Error: fmt.Errorf("couldn't create %s.\n%s", constants.COMMIT_MSG_PATH_LOCAL, hookContentMessage)}
	}

	time.Sleep(time.Second * 1)
	sub <- spinner.SpinnerResultMsg[bool]{Result: true}
}

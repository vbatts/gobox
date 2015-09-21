package common

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func getRelativeGoboxBinaryPath() (string, error) {
	callname := os.Args[0]
	// First check: Is gobox in $PATH?
	path, err := exec.LookPath(callname)
	if err == nil {
		return path, nil
	}

	// Second check: Is gobox in the current directory?
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path = filepath.Join(cwd, "gobox")
	if PathExists(path) {
		return path, nil
	}
	return "", errors.New("Could not find gobox binary")
}

func GetGoboxBinaryPath() (string, error) {
	relpath, err := getRelativeGoboxBinaryPath()
	if err != nil {
		return "", err
	}
	return filepath.Abs(relpath)
}

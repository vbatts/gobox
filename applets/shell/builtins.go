package shell

import (
	"errors"
	"os"
	"strconv"
)

type BuiltinHandler func(call []string) error

var (
	Builtins map[string]BuiltinHandler
)

func init() {
	Builtins = map[string]BuiltinHandler{
		"cd":       cd,
		"pwd":      pwd,
		"exit":     exit,
		"env":      env,
		"getenv":   getenv,
		"setenv":   setenv,
		"unsetenv": unsetenv,
		"fork":     fork,
	}
}

func pwd(call []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	println(pwd)
	return nil
}

func cd(call []string) error {
	if len(call) != 2 {
		if home := os.Getenv("HOME"); home != "" {
			return os.Chdir(home)
		}
		return os.Chdir("/")
	}
	return os.Chdir(call[1])
}

func exit(call []string) (err error) {
	code := 0
	if len(call) >= 2 {
		code, err = strconv.Atoi(call[1])
		if err != nil {
			return err
		}
	}
	os.Exit(code)
	return nil
}

func env(call []string) error {
	for _, envvar := range os.Environ() {
		println(envvar)
	}
	return nil
}

func getenv(call []string) error {
	if len(call) != 2 {
		return errors.New("`getenv <variable name>`")
	}
	println(os.Getenv(call[1]))
	return nil
}

func setenv(call []string) error {
	if len(call) != 3 {
		return errors.New("`setenv <variable name> <value>`")
	}
	return os.Setenv(call[1], call[2])
}

func unsetenv(call []string) error {
	if len(call) != 2 {
		return errors.New("`unsetenv <variable name>`")
	}
	return os.Setenv(call[1], "")
}

func fork(call []string) error {
	if len(call) < 2 {
		return errors.New("`fork <command...>`")
	}
	go execute(call[1:])
	return nil
}

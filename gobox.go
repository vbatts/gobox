package main

import (
	"flag"
	"path/filepath"

	"github.com/surma/gobox/pkg/common"
)

const (
	VERSION = "0.3.1"
)

var (
	flagSet     = flag.NewFlagSet("gobox", flag.ExitOnError)
	listFlag    = flagSet.Bool("list", false, "List applets")
	installFlag = flagSet.String("install", "", "Create symlinks for applets in given path")
	helpFlag    = flagSet.Bool("help", false, "Show help")
)

func Gobox(call []string) (err error) {
	err = flagSet.Parse(call[1:])
	if err != nil {
		return
	}

	if *listFlag {
		list()
	} else if *installFlag != "" {
		err = install(*installFlag)
	} else {
		help()
	}
	return
}

func help() {
	println("gobox [options]")
	flagSet.PrintDefaults()
	println()
	println("Version", VERSION)
	list()
}

func list() {
	println("List of compiled applets:\n")
	for name := range Applets {
		print(name, ", ")
	}
	println("")
}

func install(path string) error {
	goboxpath, err := common.GetGoboxBinaryPath()
	if err != nil {
		return err
	}
	for name := range Applets {
		// Don't overwrite the executable
		if name == "gobox" {
			continue
		}
		newpath := filepath.Join(path, name)
		err = common.ForcedSymlink(goboxpath, newpath)
		if err != nil {
			common.DumpError(err)
		}
	}
	return nil
}

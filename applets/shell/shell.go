package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/surma/gobox/pkg/common"
)

func Shell(call []string) error {
	var in *common.BufferedReader
	interactive := true
	if len(call) > 2 {
		call = call[0:1]
	}
	if len(call) == 2 {
		f, err := os.Open(call[1])
		if err != nil {
			return err
		}
		defer f.Close()
		in = common.NewBufferedReader(f)
		interactive = false
	} else {
		in = common.NewBufferedReader(os.Stdin)
	}

	var err error
	var line string
	for err == nil {
		if interactive {
			fmt.Print("> ")
		}
		line, err = in.ReadWholeLine()
		if err != nil {
			return err
		}
		if isComment(line) {
			continue
		}
		params, ce := common.Parameterize(line)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
		ce = execute(params)
		if ce != nil {
			common.DumpError(ce)
			continue
		}
	}
	return nil
}

func isComment(line string) bool {
	line = strings.TrimSpace(line)
	return strings.HasPrefix(line, "#")
}

func execute(cmdArgs []string) error {
	if len(cmdArgs) == 0 {
		return nil
	}
	if isBuiltIn(cmdArgs[0]) {
		builtin := Builtins[cmdArgs[0]]
		return builtin(cmdArgs)
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func isBuiltIn(cmd string) bool {
	_, ok := Builtins[cmd]
	return ok
}

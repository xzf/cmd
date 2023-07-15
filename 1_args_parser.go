package cmd

import (
	"errors"
	"os"
)

type argsInfo struct {
	binName     string
	subCmdSlice []string
	configSlice []kv
}

type kv struct {
	k string
	v string
}

func (cmd *CMD) parseArgs() (*argsInfo, error) {
	argSlice := os.Args
	if len(argSlice) == 0 {
		return nil, errors.New(`[4p0ahdq9u7] len(os.Args) == 0`)
	}
	result := &argsInfo{
		binName: argSlice[0],
	}
	if len(argSlice) == 1 {
		return result, nil
	}
	for i := 1; i < len(argSlice); i++ {

	}
	return result, nil
}

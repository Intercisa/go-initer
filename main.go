package main

import (
	"flag"
	"fmt"
	"os"

	"barnasipiczki.com/go-initer/constants"
	"barnasipiczki.com/go-initer/file"
	"barnasipiczki.com/go-initer/option"
	"barnasipiczki.com/go-initer/util"
)

func main() {
	defaultModule := util.GetEnv(constants.DEFAULT_MOD_ENV, constants.DEFAULT_MODULE)

	moduleInput := flag.String(constants.MODULE_INPUT_FLAG, constants.BLANK, constants.MODULE_INPUT_INFO)
	useDefault := flag.Bool(constants.DEFAULT_MODULE_FLAG, false, fmt.Sprintf(constants.DEFAULT_MODULE_INFO, defaultModule))
	flag.Parse()

	if *moduleInput == constants.BLANK {
		fmt.Println("Error occured: Please give a module name! For more info use the -i flag!")
		os.Exit(1)
	}

	var module string

	if *useDefault {
		module = defaultModule + *moduleInput
	} else {
		module = *moduleInput
	}

	err := option.ExecOptions(module)
	if err != nil {
		fmt.Println("Error during creating the moduel ", err)
		os.Exit(1)
	}
	file.WriteMainIntoFile()
}

package option

import (
	"fmt"
	"os/exec"
)

const (
	APP        = "go"
	FIRST_AG   = "mod"
	SECOND_ARG = "init"
)

func ExecOptions(module string) error {
	app := APP
	arg0 := FIRST_AG
	arg1 := SECOND_ARG
	arg2 := module

	cmd := exec.Command(app, arg0, arg1, arg2)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error during execuing command:", err)
		return err
	}
	fmt.Printf("Module: %s created", module)
	return nil
}

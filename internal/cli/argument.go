package cli

import (
	"flag"
	"fmt"
)

func attachCommandArgument(cmdInput *CommandInput) error {
	arg, err := getCommandArgument()
	if err != nil {
		return fmt.Errorf("failed to get command argument: %v", err.Error())
	}

	cmdInput.Url = arg

	return nil
}

func getCommandArgument() (string, error) {
	if len(flag.Args()) == 0 {
		return "", fmt.Errorf("no URL provided")
	}

	if len(flag.Args()) > 1 {
		return "", fmt.Errorf("too many arguments provided")
	}

	return flag.Args()[0], nil
}

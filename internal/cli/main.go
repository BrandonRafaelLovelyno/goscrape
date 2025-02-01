package cli

import (
	"fmt"
	"os"
)

func GetArguments() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("please provide the website url as the first argument")
	}

	if len(os.Args) > 2 {
		return "", fmt.Errorf("please provide only one argument as the website url")
	}

	return os.Args[1], nil
}

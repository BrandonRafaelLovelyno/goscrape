package cli

import (
	"flag"
	"fmt"
)

func GetArguments() (*Argument, error) {
	output := flag.String("o", "output.json", "JSON output file relative directory")

	flag.Parse()

	if len(flag.Args()) == 0 {
		return nil, fmt.Errorf("no URL provided")
	} else if len(flag.Args()) > 1 {
		return nil, fmt.Errorf("too many arguments provided")
	}

	return &Argument{Output: *output, Url: flag.Args()[0]}, nil
}

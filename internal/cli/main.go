package cli

import (
	"flag"
	"fmt"
)

func GetAllArguments() (*Argument, error) {
	config, err := getConfigArgument()
	if err != nil {
		fmt.Println("Executing without configuration file")
	}

	flag, err := getFlagArgument()
	if err != nil {
		return nil, err
	}

	arg := mergeArguments(flag, config)
	return arg, nil
}

func getConfigArgument() (*Config, error) {
	confDir := flag.String("c", "", "JSON configuration file relative directory")
	if *confDir == "" {
		return nil, fmt.Errorf("no configuration file provided")
	}

	config, err := extractConfigJson(*confDir)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getFlagArgument() (*FlagArgument, error) {
	outDir := flag.String("o", "output.json", "JSON output file relative directory")

	flag.Parse()

	if len(flag.Args()) == 0 {
		return nil, fmt.Errorf("no URL provided")
	}

	if len(flag.Args()) > 1 {
		return nil, fmt.Errorf("too many arguments provided")
	}

	return &FlagArgument{OutDir: *outDir, Url: flag.Args()[0]}, nil
}

func mergeArguments(flagArgument *FlagArgument, configArgument *Config) *Argument {
	arg := &Argument{
		Url:    flagArgument.Url,
		OutDir: flagArgument.OutDir,
	}

	if configArgument != nil {
		arg.WaitedElements = configArgument.WaitedElements
		arg.TargetElements = configArgument.TargetElements
		arg.Cookies = configArgument.Cookies
	}

	return arg
}

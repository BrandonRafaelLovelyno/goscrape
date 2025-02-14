package cli

import (
	"flag"
	"log"
)

func attachCommandFlag(cmdInput *CommandInput) {
	cmdFlag := getCommandFlag()

	cmdInput.OutDir = cmdFlag.OutDir

	err := attachConfig(cmdFlag.confDir, cmdInput)
	if err != nil {
		log.Printf("Failed to attach config: %v", err.Error())
	}
}

func getCommandFlag() *CommandArgument {
	confDir := flag.String("c", "", "JSON configuration file relative directory")
	outDir := flag.String("o", "output.json", "JSON output file relative directory")

	flag.Parse()

	if *confDir == "" {
		confDir = nil
		log.Printf("No configuration file provided")
	}

	return &CommandArgument{confDir: confDir, OutDir: *outDir}
}

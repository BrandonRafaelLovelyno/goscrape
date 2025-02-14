package cli

func ParseCommandInput() (*CommandInput, error) {
	var cmdInput CommandInput

	attachCommandFlag(&cmdInput)
	attachCommandArgument(&cmdInput)

	return &cmdInput, nil
}

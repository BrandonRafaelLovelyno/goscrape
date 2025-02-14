package cli

type CommandInput struct {
	OutDir          string
	Url             string
	WaitedSelectors *[]string
	TargetSelectors *[]string
	Cookies         *[]Cookie
}

type CommandArgument struct {
	OutDir  string
	confDir *string
}

type Config struct {
	waitedSelectors *[]string `json:"waitedElements"`
	targetSelectors *[]string `json:"targetElements"`
	Cookies         *[]Cookie `json:"cookie"`
}

type Cookie struct {
	Name  string
	Value string
}

package cli

type Argument struct {
	OutDir         string
	Url            string
	WaitedElements *[]string
	TargetElements *[]string
	Cookies        *[]Cookie
}

type FlagArgument struct {
	OutDir string
	Url    string
}

type Config struct {
	WaitedElements *[]string `json:"waitedElements"`
	TargetElements *[]string `json:"targetElements"`
	Cookies        *[]Cookie `json:"cookie"`
}

type Cookie struct {
	Name  string
	Value string
}

package repl

type Command struct {
	Name        string
	Description string
	Help        string
	Cmd         func(map[string]string)
	Params      []string
}

type CommandGroup struct {
	Name        string
	Description string
	Help        string
	Commands    map[string]*Command
}

func (c *CommandGroup) AddCommand(cmd *Command) {
	c.Commands[cmd.Name] = cmd
}

type Commander struct {
	Groups map[string]*CommandGroup
}

func NewCommander() *Commander {
	cmd := &Commander{
		Groups: make(map[string]*CommandGroup),
	}

	defaultGroup := NewCommandGroup(
		"default",
		"",
		"",
	)
	defaultGroup.Commands["help"] = &Command{
		Name:        "help",
		Description: "Show help",
		Cmd:         cmd.DisplayHelp,
	}
	defaultGroup.Commands["exit"] = &Command{
		Name:        "exit",
		Description: "Exit Application",
		Cmd:         nil,
	}
	cmd.Groups[defaultGroup.Name] = defaultGroup

	return cmd
}

func NewCommandGroup(name string, description string, help string) *CommandGroup {
	return &CommandGroup{
		Name:        name,
		Description: description,
		Help:        help,
		Commands:    make(map[string]*Command),
	}
}

func (c *Commander) AddCommand(cmd *Command) {
	c.Groups["default"].Commands[cmd.Name] = cmd
}

func (c *Commander) DisplayHelp(params map[string]string) {

}

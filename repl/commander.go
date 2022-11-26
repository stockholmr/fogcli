package repl

type Command struct {
	Name        string
	Description string
	Help        string
	Cmd         func(args ...string)
}

type CommandGroup struct {
	Name        string
	Description string
	Help        string
	Commands    map[string]*Command
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

	helpCommand := NewCommand(
		"help",
		"show help",
		"",
		cmd.DisplayHelp,
	)

	exitCommand := NewCommand(
		"exit",
		"exit application",
		"",
		nil,
	)

	defaultGroup.Commands[helpCommand.Name] = helpCommand
	defaultGroup.Commands[exitCommand.Name] = exitCommand
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

func NewCommand(name string, description string, help string, f func(args ...string)) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Help:        help,
		Cmd:         f,
	}
}

func (c *Commander) AddCommand(cmd *Command) {
	c.Groups["default"].Commands[cmd.Name] = cmd
}

func (c *Commander) DisplayHelp(args ...string) {

}

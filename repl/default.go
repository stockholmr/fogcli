package repl

var c *Commander

func init() {
	c = NewCommander()
}

func AddCommand(cmd *Command) {
	c.Groups["default"].Commands[cmd.Name] = cmd
}

func AddGroup(g *CommandGroup) {
	c.Groups[g.Name] = g
}

func Parse(command string) (*Command, map[string]string, error) {
	return c.Parse(command)
}

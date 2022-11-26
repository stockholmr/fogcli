package repl

var c *Commander

func init() {
	c = NewCommander()
}

func AddCommand(cmd *Command) {
	c.Groups["default"].Commands[cmd.Name] = cmd
}

func Parse(command string) (*Command, []string, error) {
	return c.Parse(command)
}

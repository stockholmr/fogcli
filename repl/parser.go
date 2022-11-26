package repl

import (
	"errors"
	"strings"
)

func (c *Commander) parseArgs(arr []string, startIndex int) ([]string, error) {
	if startIndex > len(arr)-1 {
		return nil, errors.New("invalid start index out of bounds")
	}
	newArr := make([]string, 0)
	for i := startIndex; i < len(arr); i++ {
		newArr = append(newArr, arr[i])
	}
	return newArr, nil
}

func (c *Commander) Parse(command string) (*Command, []string, error) {

	commandArr := strings.Split(command, " ")

	if len(commandArr) == 0 {
		return nil, nil, errors.New("invalid command")
	}

	var _command *Command
	var _args []string
	var err error

	firstItem := strings.ToLower(commandArr[0])

	if command, ok := c.Groups["default"].Commands[firstItem]; ok {
		// is default command
		if len(commandArr) > 1 {
			_args, err = c.parseArgs(commandArr, 1)
			if err != nil {
				return nil, nil, err
			}
		}
		_command = command
	}

	if group, ok := c.Groups[firstItem]; ok {
		// is group
		if len(commandArr) > 1 {
			secondItem := strings.ToLower(commandArr[1])
			if command, ok := group.Commands[secondItem]; ok {

				if len(commandArr) > 2 {
					_args, err = c.parseArgs(commandArr, 2)
					if err != nil {
						return nil, nil, err
					}
				}
				_command = command
			}
		}
	}

	if _command == nil {
		return nil, nil, errors.New("invalid command")
	}

	return _command, _args, nil
}

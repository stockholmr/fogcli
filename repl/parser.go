package repl

import (
	"errors"
	"strings"
)

func (c *Commander) parseParams(arr []string, startIndex int) (map[string]string, error) {
	arrayLength := len(arr)

	if startIndex > arrayLength-1 {
		return nil, errors.New("invalid start index out of bounds")
	}

	cp := make(map[string]string, 0)

	parsed := 0
	for i := startIndex; i < len(arr); i++ {
		if i != parsed {
			var paramKey, paramValue string
			paramKey = arr[i]

			if paramKey[0:1] != "-" {
				return nil, errors.New("invalid parameter. Parameters start with -")
			}

			if i+1 < arrayLength {
				paramValue = strings.TrimSpace(arr[i+1])
				parsed = i + 1
			}

			paramKey = strings.ToLower(strings.Replace(paramKey, "-", "", 1))
			cp[paramKey] = paramValue
		}
	}
	return cp, nil
}

func (c *Commander) Parse(command string) (*Command, map[string]string, error) {

	commandArr := strings.Split(command, " ")

	if len(commandArr) == 0 {
		return nil, nil, errors.New("invalid command")
	}

	var _command *Command
	var _params map[string]string
	var err error

	firstItem := strings.ToLower(commandArr[0])

	if command, ok := c.Groups["default"].Commands[firstItem]; ok {
		// is default command
		if len(commandArr) > 1 {
			_params, err = c.parseParams(commandArr, 1)
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
					_params, err = c.parseParams(commandArr, 2)
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

	return _command, _params, nil
}

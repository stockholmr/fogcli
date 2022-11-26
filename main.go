package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stockholmr/fogcli/repl"
)

func main() {

	RegisterCommands()

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		cmd, params, err := repl.Parse(reader.Text())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if cmd.Name == "exit" {
				return
			}

			// validate parmaters
			var paramsOK = true
			if len(cmd.Params) > 0 {
				for _, p := range cmd.Params {
					if _, ok := params[p]; !ok {
						fmt.Printf("missing parameter: %s\n", p)
						paramsOK = false
						break
					}
				}

			}
			// execute command
			if paramsOK {
				go cmd.Cmd(params)
			}
		}
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

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
		cmd, args, err := repl.Parse(reader.Text())
		_ = args
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if cmd.Name == "exit" {
				return
			}

			cmd.Cmd()
		}
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

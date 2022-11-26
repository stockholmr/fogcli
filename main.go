package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/stockholmr/fogcli/fog"
	"github.com/stockholmr/fogcli/repl"
	"golang.org/x/term"
)

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		cmd, args, err := repl.Parse(reader.Text())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if cmd.Name == "exit" {
				return
			}
			cmd.Cmd(args...)
		}
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

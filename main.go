package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	commands := map[string]interface{}{
		".clear":   clearScreen,
		".connect": cmd.Connect,
	}

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := cleanInput(reader.Text())

		if command, exists := commands[text]; exists {
			// Call a hardcoded function
			command.(func())()
		} else if strings.EqualFold(".exit", text) {
			// Close the program on the exit command
			return
		} else {
			// Pass the command to the parser
			//handleCmd(text)
		}
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

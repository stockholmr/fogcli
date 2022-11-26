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

func RegisterCommands() {

	repl.AddCommand(&repl.Command{
		Name:        "clear",
		Description: "clear console window",
		Cmd: func(args ...string) {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
	}) // end clear

	repl.AddCommand(&repl.Command{
		Name:        "connect",
		Description: "Connect to fogserver database",
		Cmd: func(args ...string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Fogserver: ")
			hostname, _ := reader.ReadString('\n')
			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
			fmt.Print("\n")
			fmt.Print("Database: ")
			database, _ := reader.ReadString('\n')

			s_hostname := strings.TrimSpace(hostname)
			s_username := strings.TrimSpace(username)
			s_password := strings.TrimSpace(string(bytePassword))
			s_database := strings.TrimSpace(database)

			fog.Connect(s_hostname, s_username, s_password, s_database)
		},
	})// end connect

} // end RegisterCommands

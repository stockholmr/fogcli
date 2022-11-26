package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/stockholmr/fogcli/fog"
	"golang.org/x/term"
    "github.com/stockholmr/fogcli/repl"
)

repl.AddCommand(repl.NewCommand(
    "connect",
    "Connect to fogserver database",
    "",
    func(args ...string) {
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
))

func Connect() {
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
}

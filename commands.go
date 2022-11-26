package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/stockholmr/fogcli/fog"
	"github.com/stockholmr/fogcli/repl"
)

func RegisterCommands() {

	repl.AddCommand(&repl.Command{
		Name:        "clear",
		Description: "clear console window",
		Cmd: func(params map[string]string) {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
	}) // end clear

	repl.AddCommand(&repl.Command{
		Name:        "connect",
		Description: "Connect to fogserver database",
		Params:      []string{"server", "username", "password", "database"},
		Cmd: func(params map[string]string) {
			fog.Connect(params["server"], params["username"], params["password"], params["database"])
		},
	}) // end connect

	hostGroup := repl.NewCommandGroup("host", "", "")
	hostGroup.AddCommand(&repl.Command{
		Name:        "list",
		Description: "List Hosts",
		Cmd: func(params map[string]string) {
			hosts, err := fog.Hosts_List()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Print(hosts)
		},
	}) // end hosts.list

	repl.AddGroup(hostGroup)

} // end RegisterCommands

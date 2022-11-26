package cmd

import (
	"os"
	"os/exec"
    "github.com/stockholmr/fogcli/repl"
)

repl.AddCommand(repl.NewCommand(
    "clear",
    "",
    "",
    func(args ...string) {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    },
))
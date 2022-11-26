package cmd

import (
	"fmt"

	"github.com/stockholmr/fogcli/fog"
	"github.com/stockholmr/logging"
)

func Hosts_List() {
	hosts, err := fog.Hosts_List()
	if err != nil {
		logging.Error(err.Error())
	}

	fmt.Print(hosts)
}

func Hosts_Search() {
	hosts, err := fog.Hosts_Find()
}

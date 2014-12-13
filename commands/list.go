package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all hosts",
	Long:  `Shows you what hosts you have configured`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
		for _, host := range hosts.Hosts {
			fmt.Println(host.Name)
			fmt.Println(host.Config)
			fmt.Println("")
		}
	},
}

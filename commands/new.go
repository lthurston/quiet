package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"

	"github.com/spf13/cobra"
)

var newFrom string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Appends new host using other host as template",
	Long:  `New appends a new host to your SSH configuration based on other host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("from: " + newFrom)

		from := config.GetConfigNewFrom()
		if len(newFrom) > 0 {
			from = newFrom
		}

		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
	},
}

package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"

	"github.com/spf13/cobra"
)

var newFrom, newName string
var newSkipInteractive bool

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Appends new host using other host as template",
	Long:  `New appends a new host to your SSH configuration based on other host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("from: ", newFrom)
		fmt.Println("new: ", newName)
		fmt.Println("skip-interactive: ", newSkipInteractive)

		from := config.GetConfigNewFrom()
		if len(newFrom) > 0 {
			from = newFrom
		}

		if len(newName) == 0 {
			// Cheapy way to make new name
			newName = from + "0"
		}

		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
		if host, found := hosts.FindHostByName(from); found {
			host.Name = newName
			newHostSnippet := host.RenderSnippet()
			fmt.Println(newHostSnippet)
		} else {
			fmt.Println("Couldn't find host or empty: \"" + from + "\"")
		}
	},
}

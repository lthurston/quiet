package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"
	"github.com/spf13/cobra"
)

var warpCmd = &cobra.Command{
	Use:   "warp",
	Short: "Generates warp config",
	Long:  `Creates a warp config file based on current hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
		for _, host := range hosts.HostPositions {
			fmt.Println(host.Name())
		}
	},
}

package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/host"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds hosts based on search arguments",
	Long:  `Finds hosts based on search arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := host.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())

		for _,host := range hosts.HostPositions {
			if(host.ContainsStrings(args)) {
				fmt.Println(host)
			}
		}
	},
}
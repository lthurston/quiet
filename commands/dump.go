package commands

import (
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/host"
	"github.com/spf13/cobra"
	"fmt"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps everything Quiet knows",
	Long:  `Outputs all hosts Quiet knows about`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := host.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
		for _, host := range hosts.HostPositions {
			fmt.Print(host)
		}
	},
}

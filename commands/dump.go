package commands

import (
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/host"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps everything Quiet knows",
	Long:  `Outputs any config and all hosts it knows about`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := host.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
	},
}

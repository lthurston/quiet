package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/host"
	"github.com/spf13/cobra"
	"regexp"
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
				hostString := host.String()
				for _, arg := range args {
					escArg := regexp.QuoteMeta(arg)
					re := regexp.MustCompile("(?i)(" + escArg +")")
					hostString = re.ReplaceAllString(hostString,"\x1b[7m $1 \x1b[0m")
				}
				fmt.Println(hostString)
			}
		}
	},
}
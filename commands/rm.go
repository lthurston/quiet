package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/writer"
	"github.com/lthurston/quiet/splicer"
	"bytes"
	"os"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a host",
	Long:  `Deletes one or more hosts from config`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := host.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())

		if host, found := hosts.FindHostByName(args[0]); found {
			var buffer bytes.Buffer
			f, err := os.Open(config.GetConfigFile())
			if err != nil {
				panic(err)
			}

			defer f.Close()

			// TODO: Splicer line numbering is weird. Ideally, we'd put the
			// StartLine() as the splice from and EndLine() as splice too, any everything
			// would taste like candy canes.
			splicer.SpliceInto(host.StartLine() - 1, host.EndLine(), "", f, &buffer)
			writer.Replace(buffer.String())
		} else {
			fmt.Println("Couldn't find host or empty: \"" + args[0] + "\"")
		}
	},
}

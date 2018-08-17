package commands

import (
	"github.com/spf13/cobra"
	"github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/splicer"
	"github.com/lthurston/quiet/writer"
	"bytes"
	"os"
)

var force bool

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Reformats the ssh config file",
	Long:  `Iterates the ssh config file, replacing each entry with a reformatted one`,
}

func init() {
	tidyCmd.Flags().BoolVarP(&force, "force", "f", false, "Forces the overwriting of the ssh config file")
	tidyCmd.Run = tidy
}

func tidy(cmd *cobra.Command, args []string) {
	hosts := host.HostsCollection{}
	hosts.ReadFromFile(config.GetConfigFile())

	var buffer bytes.Buffer
	f, err := os.Open(config.GetConfigFile())
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//splicer.SpliceInto(host.StartLine(), host.EndLine(), "", f, &buffer)
	//writer.Replace(buffer.String())

	// iterate the hosts
	for _, host := range hosts.HostPositions {
		// render the host
		splicer.SpliceInto(host.StartLine(), host.EndLine(), host.String(), f, &buffer)
		writer.Replace(buffer.String())
	}
}

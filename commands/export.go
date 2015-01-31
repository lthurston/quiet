package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/util"
	"strings"
	"io/ioutil"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports config snippet and keys for sharing",
	Long:  `Exports one or more host snippets and associated keys for easy distribution`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := host.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())

		if host, found := hosts.FindHostByName(args[0]); found {
			fmt.Println(host)

			fields := strings.Split(config.GetConfigExportFilenameOptions(), ",")
			for _, field := range fields {
				filename := host.GetOptionArgument(strings.TrimSpace(field))

				if(len(filename) > 0) {
					if fileContents, err := ioutil.ReadFile(util.Detilde(filename)); err == nil {
						fmt.Println(filename, "\n")
						fmt.Println(string(fileContents))

						// If there's a file by the same name with .pub at the end, show it too
						filenamePub := filename + ".pub"
						if fileContentsPub, err := ioutil.ReadFile(util.Detilde(filenamePub)); err == nil {
							fmt.Println(filenamePub, "\n")
							fmt.Println(string(fileContentsPub))
						}

					} else {
						fmt.Println("Couldn't find: ", filename)
					}
				}
			}

		} else {
			fmt.Println("Couldn't find host or empty: \"" + args[0] + "\"")
		}
	},
}

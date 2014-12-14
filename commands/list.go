package commands

import (
	"fmt"
	"strings"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all hosts",
	Long:  `Shows you what hosts you have configured`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())
		for _, host := range hosts.Hosts {
			configFields := formatHostConfigFields(config.GetConfigListFields(), host.Config)
			fmt.Printf("%-20s %s", host.Name, configFields)
			fmt.Println("")
		}
	},
}

func formatHostConfigFields(fieldsString string, configs map[string]string) string {
	out := ""
	fields := strings.Split(fieldsString, ",")
	for key, value := range configs {
		for _, field := range fields {
			if key == strings.TrimSpace(field) {
				out = out + field + ": " + value + "  "
			}
		}
	}
	return out
}

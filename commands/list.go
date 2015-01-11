package commands

import (
	"fmt"
	"strconv"
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

		widths := getColumnWidths(hosts)
		fields := strings.Split(config.GetConfigListFields(), ",")
		for _, host := range hosts.HostPositions {
			fmt.Printf("%-"+strconv.Itoa(widths["Name"]+1)+"s", host.Name())
			for _, field := range fields {
				fieldName := strings.TrimSpace(field)
				fmt.Printf("%-"+strconv.Itoa(widths[fieldName]+1)+"s", host.Config()[fieldName])
			}
			fmt.Println("")
		}
	},
}

func getColumnWidths(hosts parser.HostsCollection) map[string]int {
	widths := make(map[string]int)
	for _, host := range hosts.HostPositions {
		if length := len([]rune(host.Name())); length > widths["Name"] {
			widths["Name"] = length
		}
		for key, value := range host.Config() {
			if length := len([]rune(value)); length > widths[key] {
				widths[key] = length
			}
		}
	}
	return widths
}

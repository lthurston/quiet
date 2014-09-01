package commands

import (
	"fmt"

	"github.com/lthurston/quiet/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Lists / modifies configuration",
	Long:  `Shows all configuration values if no arguments are provided, otherwise modifies with argument key/value pairs`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Parse("")
		fmt.Println("file: " + config.GetSSHFile())
		fmt.Println("strategy: " + config.GetStrategy())
	},
}

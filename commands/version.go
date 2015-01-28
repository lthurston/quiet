package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const QuietVersion = "0.2 alpha"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Quiet",
	Long:  `All software has versions. This is Quiet's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Quiet ssh shusher " + QuietVersion)
	},
}

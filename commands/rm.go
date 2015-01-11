package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a host",
	Long:  `Deletes one or more hosts from config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Implement me!")
	},
}

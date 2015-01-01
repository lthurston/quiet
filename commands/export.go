package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports config snippet and keys for sharing",
	Long:  `Exports one or more host snippets and associated keys for easy distribution`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
	},
}

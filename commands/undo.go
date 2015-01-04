package commands

import (
	"github.com/lthurston/quiet/writer"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last quiet action",
	Long:  `Swaps the current config with the last backup`,
	Run: func(cmd *cobra.Command, args []string) {
		writer.Undo()
	},
}

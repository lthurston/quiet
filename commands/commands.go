package commands

import (
	"github.com/spf13/cobra"
)

var QuietCmd = &cobra.Command{
	Use:   "quiet",
	Short: "Quiet helps keep down the noise by managing your ssh config file",
	Long: `Quiet allows you to quickly add _new_ ssh Hosts, _copy_ existing
        ones, _delete_ them, and maybe some other stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	addCommands()
	QuietCmd.Execute()
}

func addCommands() {
	QuietCmd.AddCommand(versionCmd)
	QuietCmd.AddCommand(configCmd)
	QuietCmd.AddCommand(dumpCmd)
	QuietCmd.AddCommand(listCmd)
	QuietCmd.AddCommand(exportCmd)
	QuietCmd.AddCommand(undoCmd)
	QuietCmd.AddCommand(warpCmd)
	QuietCmd.AddCommand(rmCmd)
	QuietCmd.AddCommand(newCmd)
	QuietCmd.AddCommand(findCmd)
	QuietCmd.AddCommand(tinyCmd)
}

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

	// find a better spot to put the flags
	newCmd.Flags().StringVarP(&newFrom, "from", "f", "", "host to use as template")
	newCmd.Flags().StringVarP(&newName, "name", "n", "", "new host name")
	newCmd.Flags().BoolVarP(&newSkipInteractive, "skip-interactive", "s", false, "just copy; don't allow interactive")
	QuietCmd.AddCommand(newCmd)
}

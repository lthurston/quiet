package commands

import "github.com/spf13/cobra"

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps everything Quiet knows",
	Long:  `Outputs any config and all hosts it knows about`,
	Run: func(cmd *cobra.Command, args []string) {
		// Does ssh config file exist?

		// If yes, output the values

		// Does an SSH config file exist in the expected location?

		// If yes, output the hosts

	},
}

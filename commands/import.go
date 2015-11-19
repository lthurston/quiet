package commands

import (
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import < something",
	Short: "Imports config snippet and keys",
	Long:  `Imports a host snippet and associated keys`,
}

func init() {
	importCmd.Run = importCmdRun
}

func importCmdRun(cmd *cobra.Command, args []string) {
	// Steps
	// * open export tarball
	// * confirm presence of all associated key files (prompt to continue)
	// * check if hostname already exists here, do something if so (prompt for rename)
	// * check if private key content already exists in any key file, if so, prompt to consolidate
	// * check if key filename exists (if so, compare contents, then prompt to rename)
	// * with updated host, write config, write update key files
}

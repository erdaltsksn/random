package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// stateCmd represents the state command.
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Generate a random state",
	Long:  `Generate a random state using the pre-defined list.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.State())
	},
}

func init() {
	rootCmd.AddCommand(stateCmd)
}

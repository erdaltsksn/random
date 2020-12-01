package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// tcCmd represents the tc command.
var tcCmd = &cobra.Command{
	Use:     "tc",
	Aliases: []string{"TC"},
	Short:   "Generate a random TC",
	Long:    `Generate a random TC (Turkish Citizen Number)`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.TC())
	},
}

func init() {
	RootCmd.AddCommand(tcCmd)
}

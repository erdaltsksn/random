package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// nameLastCmd represents the name command.
var nameLastCmd = &cobra.Command{
	Use:   "last",
	Short: "Generate a random name",
	Long:  `Generate a random name using the country specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Lastname(country))
	},
}

func init() {
	nameCmd.AddCommand(nameLastCmd)
}

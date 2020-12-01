package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// nameFirstCmd represents the name command.
var nameFirstCmd = &cobra.Command{
	Use:   "first",
	Short: "Generate a random name",
	Long:  `Generate a random name using gender and country specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Firstname(gender, country))
	},
}

func init() {
	nameCmd.AddCommand(nameFirstCmd)
}

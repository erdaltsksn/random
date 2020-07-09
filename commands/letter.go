package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// letterCmd represents the letter command.
var letterCmd = &cobra.Command{
	Use:   "letter",
	Short: "Generates a random letter",
	Long:  `Generates a random latin letter in the range of a-Z and A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Letter())
	},
}

func init() {
	rootCmd.AddCommand(letterCmd)
}

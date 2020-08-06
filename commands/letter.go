package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// letterCmd represents the letter command.
var letterCmd = &cobra.Command{
	Use:   "letter",
	Short: "Generate a random letter",
	Long:  `Generate a random latin letter in the range of a-Z and A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Letter())
	},
}

func init() {
	rootCmd.AddCommand(letterCmd)
}

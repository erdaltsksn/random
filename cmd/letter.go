package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// letterCmd represents the letter command
var letterCmd = &cobra.Command{
	Use:   "letter",
	Short: "Generates a random letter",
	Long:  `Generates a random latin letter in the range of a-Z and A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		Output(random.Letter())
	},
}

func init() {
	rootCmd.AddCommand(letterCmd)
}

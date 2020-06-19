package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "Generates a random uppercase letter",
	Long:  `Generates a random uppercase latin letter in the range of A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		Output(random.LetterUpperCase())
	},
}

func init() {
	letterCmd.AddCommand(upperCmd)
}

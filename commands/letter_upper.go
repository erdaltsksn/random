package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// letterUpperCmd represents the upper command.
var letterUpperCmd = &cobra.Command{
	Use:   "upper",
	Short: "Generates a random uppercase letter",
	Long:  `Generates a random uppercase latin letter in the range of A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LetterUpperCase())
	},
}

func init() {
	letterCmd.AddCommand(letterUpperCmd)
}

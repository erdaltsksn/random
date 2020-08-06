package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// letterLowerCmd represents the lower command.
var letterLowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "Generate a random lowercase letter",
	Long:  `Generate a random lowercase latin letter in the range of a-z.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LetterLowerCase())
	},
}

func init() {
	letterCmd.AddCommand(letterLowerCmd)
}

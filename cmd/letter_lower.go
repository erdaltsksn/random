package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "Generates a random lowercase letter",
	Long:  `Generates a random lowercase latin letter in the range of a-z.`,
	Run: func(cmd *cobra.Command, args []string) {
		Output(random.LetterLowerCase())
	},
}

func init() {
	letterCmd.AddCommand(lowerCmd)
}

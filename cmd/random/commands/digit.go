package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// digitCmd represents the digit command.
var digitCmd = &cobra.Command{
	Use:   "digit",
	Short: "Generate a random digit",
	Long: `Generate a random digits (as one-digit-numerals). A digit is a number
from 0 to 9.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Digit())
	},
}

func init() {
	RootCmd.AddCommand(digitCmd)
}

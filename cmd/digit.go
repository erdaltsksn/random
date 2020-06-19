package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// digitCmd represents the digit command
var digitCmd = &cobra.Command{
	Use:   "digit",
	Short: "Generates a random digit",
	Long: `Generates a random digits (as one-digit-numerals). A digit is a
number from 0 to 9.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Digit())
	},
}

func init() {
	rootCmd.AddCommand(digitCmd)
}

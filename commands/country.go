package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// countryCmd represents the country command.
var countryCmd = &cobra.Command{
	Use:   "country",
	Short: "Generates a random country name",
	Long:  `Generates a random country name using the pre-defined list.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Country())
	},
}

func init() {
	rootCmd.AddCommand(countryCmd)
}

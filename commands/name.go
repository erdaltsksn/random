package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

var (
	gender  string
	country string
)

// nameCmd represents the name command.
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Generate a random name",
	Long:  `Generate a random name using gender and country specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Name(gender, country))
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.
	nameCmd.Flags().StringVarP(&gender, "gender", "g", "",
		`Which gender should be used. Avaible genders:
- Male
- Female`,
	)
	nameCmd.Flags().StringVarP(&country, "country", "c", "",
		`Which country should be used. Available countries:
- USA
- Germany
- Estonia
- Russia
- Turkey`,
	)
}

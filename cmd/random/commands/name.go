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
	Short: "Generate a random full name",
	Long:  `Generate a random full name using gender and country specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Fullname(gender, country))
	},
}

func init() {
	RootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.
	nameCmd.PersistentFlags().StringVarP(&gender, "gender", "g", "",
		`Which gender should be used. Available genders:
- Male
- Female`,
	)
	nameCmd.PersistentFlags().StringVarP(&country, "country", "c", "",
		`Which country should be used. Available countries:
- USA
- Germany
- Estonia
- Russia
- Turkey`,
	)
}

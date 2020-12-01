package commands

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

var listAuthors bool

// authorCmd represents the author command.
var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Generate a random author",
	Long:  `Generate a random author or list all available authors`,
	Run: func(cmd *cobra.Command, args []string) {
		if listAuthors {
			var data string
			for _, v := range random.QuoteList {
				data += v.Author + "\n"
			}
			printOutput(strings.Trim(data, "\n"))
			os.Exit(0)
		}

		printOutput(random.Quote("").Author)
	},
}

func init() {
	RootCmd.AddCommand(authorCmd)

	// Here you will define your flags and configuration settings.
	authorCmd.Flags().BoolVar(&listAuthors, "all", false,
		`Option to list all authors`)
}

package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

var author string

// quoteCmd represents the quote command.
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Generate a random quote",
	Long:  `Generate a random quote based on author name if supplied.`,
	Run: func(cmd *cobra.Command, args []string) {
		quote := random.Quote(author)
		printOutput(fmt.Sprintf("%s\n- %s", quote.Text, quote.Author))
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	// Here you will define your flags and configuration settings.
	quoteCmd.Flags().StringVarP(&author, "author", "a", "",
		`Full name of the author. To list Available authors run the following:
$random author --all`)
}

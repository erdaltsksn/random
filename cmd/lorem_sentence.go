package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// sentenceCmd represents the sentence command
var sentenceCmd = &cobra.Command{
	Use:     "sentence",
	Aliases: []string{"s"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the sentence count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremSentence(count))
	},
}

func init() {
	loremCmd.AddCommand(sentenceCmd)
}

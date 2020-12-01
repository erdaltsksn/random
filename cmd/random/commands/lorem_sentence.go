package commands

import (
	"github.com/erdaltsksn/random"
	"github.com/spf13/cobra"
)

// loremSentenceCmd represents the sentence command.
var loremSentenceCmd = &cobra.Command{
	Use:     "sentence",
	Aliases: []string{"s"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the sentence count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremSentence(count))
	},
}

func init() {
	loremCmd.AddCommand(loremSentenceCmd)
}

package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// loremParagraphCmd represents the paragraph command.
var loremParagraphCmd = &cobra.Command{
	Use:     "paragraph",
	Aliases: []string{"p"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the paragraph count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremParagraph(count))
	},
}

func init() {
	loremCmd.AddCommand(loremParagraphCmd)
}

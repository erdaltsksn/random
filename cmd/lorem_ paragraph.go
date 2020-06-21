package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// paragraphCmd represents the paragraph command
var paragraphCmd = &cobra.Command{
	Use:     "paragraph",
	Aliases: []string{"p"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the paragraph count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremParagraph(count))
	},
}

func init() {
	loremCmd.AddCommand(paragraphCmd)
}

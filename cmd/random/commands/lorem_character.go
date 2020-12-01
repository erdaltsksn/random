package commands

import (
	"github.com/erdaltsksn/random"
	"github.com/spf13/cobra"
)

// loremCharacterCmd represents the character command.
var loremCharacterCmd = &cobra.Command{
	Use:     "character",
	Aliases: []string{"c"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the character count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremCharacter(count))
	},
}

func init() {
	loremCmd.AddCommand(loremCharacterCmd)
}

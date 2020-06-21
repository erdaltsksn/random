package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:     "character",
	Aliases: []string{"c"},
	Short:   "Generate random text",
	Long:    `Generate random lorem text according to the character count.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.LoremCharacter(count))
	},
}

func init() {
	loremCmd.AddCommand(characterCmd)
}

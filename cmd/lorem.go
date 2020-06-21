package cmd

import (
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

var count int

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "lorem",
	Short: "Generate a dummy text",
	Long: `Generate a random lorem text commonly used in the graphic, print, and
publishing industries for previewing layouts and visual mockups.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Lorem())
	},
}

func init() {
	rootCmd.AddCommand(loremCmd)

	// Here you will define your flags and configuration settings.
	loremCmd.PersistentFlags().IntVarP(&count, "count", "c", 1,
		`How many items (paragraphs, sentences or characters) in lorem should be
generated.`)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"

	random "github.com/erdaltsksn/random/v1"
)

var (
	lower bool
	upper bool
)

// letterCmd represents the letter command
var letterCmd = &cobra.Command{
	Use:   "letter",
	Short: "Generates a random letter",
	Long:  `Generates a random latin letter in the range of a-Z and A-Z.`,
	Run: func(cmd *cobra.Command, args []string) {
		if lower && upper {
			color.Danger.Println("You cannot use both --lower and --upper params")
			color.Info.Prompt(`Use either --lower or --upper as a parameter`)
			os.Exit(1)
		}

		fmt.Println(random.Letter(lower, upper))
	},
}

func init() {
	rootCmd.AddCommand(letterCmd)

	// Here you will define your flags and configuration settings.
	letterCmd.PersistentFlags().BoolVarP(&lower, "lower", "l", false,
		`output lowercase letter.`)
	letterCmd.PersistentFlags().BoolVarP(&upper, "upper", "u", false,
		`output uppercase letter.`)
}

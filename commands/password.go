package commands

import (
	"github.com/erdaltsksn/random"
	"github.com/spf13/cobra"
)

var length int

// passwordCmd represents the password command.
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate password",
	Long:  `Generate a random, secure password`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.Password(length))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	// Here you will define your flags and configuration settings.
	passwordCmd.Flags().IntVarP(&length, "length", "l", 32,
		"Total length of generated password")
}

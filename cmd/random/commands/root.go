package commands

import (
	"github.com/spf13/cobra"
)

var unattended bool

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "random",
	Short: "This app helps you generate random data",
	Long: `Generate/fake random data such as number, string, password, email,
address, hex color etc.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	// Here you will define your flags and configuration settings.
	RootCmd.PersistentFlags().BoolVar(&unattended, "unattended", false,
		`Do not copy the value to the clipboard`)
}

package cmd

import (
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "random",
	Short: "This app helps you generate random data",
	Long: `Generate/fake random data such as number, string, password, email,
address, hex color etc.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Danger.Prompt(err.Error())
		os.Exit(1)
	}
}

// GetRootCmd returns the instance of root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}

package cmd

import (
	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
)

var version = "unknown"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the app's version",
	Long: `Prints the app's version. This app uses semver and the version value
is automatically generated while building.`,
	Run: func(cmd *cobra.Command, args []string) {
		cui.Info(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
)

var version string

// SetVersion set the application version for consumption in the output of the command.
func SetVersion(v string) {
	version = v
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the app's version",
	Long: `Prints the app's version. This app uses semver and the version value
automatically generated while building.`,
	Run: func(cmd *cobra.Command, args []string) {
		cui.Info(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

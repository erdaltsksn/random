package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

// userAgentCmd represents the userAgent command.
var userAgentCmd = &cobra.Command{
	Use:     "user-agent",
	Aliases: []string{"useragent", "ua", "agent", "browser"},
	Short:   "Generate a random user-agent",
	Long:    `Generate a random user-agent using the pre-defined list.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.UserAgent())

	},
}

func init() {
	RootCmd.AddCommand(userAgentCmd)
}

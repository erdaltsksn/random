package commands

import (
	"github.com/erdaltsksn/cui"
)

func init() {
	RootCmd.AddCommand(cui.VersionCmd)
}

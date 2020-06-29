package main

import (
	"github.com/erdaltsksn/random/cmd"
)

var version = "unknown"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}

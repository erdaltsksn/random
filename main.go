package main

import (
	"github.com/erdaltsksn/random/cmd"
)

var version = "undefined"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}

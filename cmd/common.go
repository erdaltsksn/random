package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/erdaltsksn/cui"
)

// printOutput is used to print the result and copy it into system clipboard
func printOutput(o interface{}) {
	output := fmt.Sprint(o)
	fmt.Println(output)

	if err := toClipboard(output); err != nil {
		cui.Error("Error while copying value to system clipboard", err)
	}

	cui.Success("Copied")
}

func toClipboard(output string) error {
	var copyCmd *exec.Cmd

	if runtime.GOOS == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}
	if runtime.GOOS == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}
	if runtime.GOOS == "windows" {
		copyCmd = exec.Command("clip")
	}

	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}

	if _, err := in.Write([]byte(output)); err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return copyCmd.Wait()
}

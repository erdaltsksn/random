package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gookit/color"
)

// Output is used to print the result and copy it into system clipboard
func Output(o interface{}) {
	output := fmt.Sprint(o)
	fmt.Println(output)

	if err := toClipboard(fmt.Sprint(output)); err != nil {
		color.Danger.Println("Error while copying value to system clipboard")
		color.Warn.Prompt(err.Error())
		os.Exit(1)
	}

	color.Info.Println("√ Copied")
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

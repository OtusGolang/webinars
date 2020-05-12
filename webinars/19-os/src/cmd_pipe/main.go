package main

import (
	"os"
	"os/exec"
)

func main() {
	lsCmd := exec.Command("ls")
	wcCmd := exec.Command("wc", "-l")

	pipe, _ := lsCmd.StdoutPipe()
	wcCmd.Stdin = pipe
	wcCmd.Stdout = os.Stdout

	_ = lsCmd.Start()
	_ = wcCmd.Start()
	_ = lsCmd.Wait()
	_ = wcCmd.Wait()
}

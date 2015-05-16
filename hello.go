package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Printf("hello, world\n")
	cmd := NewCmd("ps -a")
	fmt.Printf("Command output %s", cmd)
}

// NewCmd is a convenience function to construct `*exec.Cmd` from string input.
func NewCmd(command string) *exec.Cmd {
	wd, _ := os.Getwd()

	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:len(parts)]

	cmd := exec.Command(head, parts...)
	cmd.Dir = wd
	cmd.Env = os.Environ()

	return cmd
}


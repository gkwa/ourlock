package common

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
	"github.com/mitchellh/go-homedir"
)

var (
	DryRun      bool
	Directories []string
)

func Run(rgArgs []string, directories []string) {
	expandedDirectories := expandTilde(directories)

	if len(expandedDirectories) == 0 {
		expandedDirectories = expandTilde(Directories)
	}

	allArgs := append(rgArgs, expandedDirectories...)

	if DryRun {
		fmt.Printf("rg %s\n", shellquote.Join(allArgs...))
		return
	}

	rgCmd := exec.Command("rg", allArgs...)

	var stdout, stderr bytes.Buffer
	rgCmd.Stdout = &stdout
	rgCmd.Stderr = &stderr

	err := rgCmd.Run()
	if err != nil {
		fmt.Println("Error running rg command:", err)
		fmt.Println("Stdout:", stdout.String())
		fmt.Println("Stderr:", stderr.String())
		os.Exit(1)
	} else {
		fmt.Print(stdout.String())
	}
}

func expandTilde(directories []string) []string {
	var expandedDirectories []string
	for _, dir := range directories {
		expanded, err := homedir.Expand(dir)
		if err != nil {
			fmt.Printf("Error expanding directory '%s': %v\n", dir, err)
			continue
		}
		expandedDirectories = append(expandedDirectories, expanded)
	}
	return expandedDirectories
}

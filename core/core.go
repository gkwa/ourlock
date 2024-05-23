package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kballard/go-shellquote"
	"github.com/mitchellh/go-homedir"
)

func Run(searchStrings []string, directories []string, dryRun bool) {
	expandedDirectories := expandTilde(directories)

	regexPattern := buildRegexPattern(searchStrings)

	rgArgs := []string{
		regexPattern,
		"--ignore-case",
		"--glob-case-insensitive",
		"--multiline-dotall",
		"--color=always",
		"--glob=*.{md,txt,org}",
		"--context=20",
	}
	rgArgs = append(rgArgs, expandedDirectories...)

	if dryRun {
		fmt.Printf("rg %s\n", shellquote.Join(rgArgs...))
		return
	}

	rgCmd := exec.Command("rg", rgArgs...)
	rgCmd.Stdout = os.Stdout
	rgCmd.Stderr = os.Stderr

	err := rgCmd.Run()
	if err != nil {
		fmt.Println("Error running rg command:", err)
		os.Exit(1)
	}
}

func buildRegexPattern(searchStrings []string) string {
	if len(searchStrings) == 1 {
		return searchStrings[0]
	}

	var patterns []string
	for i := 0; i < len(searchStrings); i++ {
		for j := 0; j < len(searchStrings); j++ {
			if i != j {
				patterns = append(patterns, fmt.Sprintf("%s(?s:.*?)%s", searchStrings[i], searchStrings[j]))
			}
		}
	}
	return strings.Join(patterns, "|")
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

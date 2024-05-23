package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func Run(searchStrings []string) {
	directories := []string{
		"/Users/mtm/Documents/Obsidian\\ Vault/",
		"~/pdev/taylormonacelli/notes/",
		"~/pdev/michell/",
	}

	expandedDirectories := expandTilde(directories)

	regexPattern := buildRegexPattern(searchStrings)

	rgCmd := exec.Command("rg", "--multiline-dotall", "--color=always", "--context=20", regexPattern, "--glob=*.{md,txt,org}", strings.Join(expandedDirectories, " "))

	rgCmd.Stdout = os.Stdout
	rgCmd.Stderr = os.Stderr

	err := rgCmd.Run()
	if err != nil {
		fmt.Println("Error running rg command:", err)
		os.Exit(1)
	}
}

func buildRegexPattern(searchStrings []string) string {
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

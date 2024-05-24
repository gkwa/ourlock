package existancesearch

import (
	"fmt"
	"strings"

	"github.com/gkwa/ourlock/common"
)

func Run(searchStrings []string, directories []string) {
	regexPattern := buildRegexPattern(searchStrings)

	rgArgs := []string{
		regexPattern,
		"--ignore-case",
		"--glob-case-insensitive",
		"--multiline",
		"--multiline-dotall",
		"--color=always",
		"--glob=*.{md,txt,org}",
		"--files-with-matches",
	}

	common.Run(rgArgs, directories)
}

func buildRegexPattern(searchStrings []string) string {
	var patterns []string
	permutations := GeneratePermutations(searchStrings)

	if common.DryRun {
		fmt.Printf("%d Permutations:\n", len(permutations))
		for _, perm := range permutations {
			fmt.Println(perm)
		}
	}

	for _, perm := range permutations {
		patterns = append(patterns, strings.Join(perm, ".+?"))
	}
	return strings.Join(patterns, "|")
}

func GeneratePermutations(slice []string) [][]string {
	if len(slice) == 1 {
		return [][]string{slice}
	}

	var result [][]string
	for i := range slice {
		rest := make([]string, len(slice)-1)
		copy(rest[:i], slice[:i])
		copy(rest[i:], slice[i+1:])
		for _, p := range GeneratePermutations(rest) {
			result = append(result, append([]string{slice[i]}, p...))
		}
	}
	return result
}

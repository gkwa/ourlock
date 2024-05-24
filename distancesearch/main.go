package distancesearch

import (
	"fmt"

	"github.com/gkwa/ourlock/common"
)

func Run(searchStrings []string, directories []string, distance int) {
	regexPattern := buildRegexPattern(searchStrings, distance)

	rgArgs := []string{
		regexPattern,
		"--ignore-case",
		"--glob-case-insensitive",
		"--multiline",
		"--multiline-dotall",
		"--color=always",
		"--glob=*.{md,txt,org}",
		"--context=20",
	}

	common.Run(rgArgs, directories)
}

func buildRegexPattern(searchStrings []string, chars int) string {
	return fmt.Sprintf("%s.{0,%d}?%s", searchStrings[0], chars, searchStrings[1])
}

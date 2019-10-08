package marker

import (
	"fmt"

	"github.com/fatih/color"
)

// Mark marks the patterns that returned from MatcherFunc with colors in given string
func Mark(str string, matcherFunc MatcherFunc, c *color.Color) string {
	match := matcherFunc(str)
	patterns := match.Patterns
	colorizeStrings(patterns, c)
	args := convertToInterfaceSlice(patterns)
	return fmt.Sprintf(match.Template, args...)
}

// MarkMany marks each set of patterns returns by a variable number of MatcherFunc with color in given string
func MarkMany(str string, c *color.Color, matcherFuncs ...MatcherFunc) string {
	for _, matcherFunc := range matcherFuncs {
		str = Mark(str, matcherFunc, c)
	}
	return str
}

func colorizeStrings(strs []string, c *color.Color) {
	for i := range strs {
		strs[i] = c.Sprintf("%s", strs[i])
	}
}

func convertToInterfaceSlice(patterns []string) []interface{} {
	args := make([]interface{}, len(patterns))
	for i := range patterns {
		args[i] = patterns[i]
	}
	return args
}

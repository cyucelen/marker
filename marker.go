package marker

import (
	"fmt"
	"github.com/fatih/color"
)

// Mark marks the patterns that returned from MatcherFunc with colors in given string
func Mark(str string, matcherFunc MatcherFunc, c *color.Color) string {
	match := matcherFunc(str)
	colorizedPatterns := colorizeStrings(match.Patterns, c)
	args := convertToInterfaceSlice(colorizedPatterns)
	return fmt.Sprintf(match.Template, args...)
}

func colorizeStrings(strs []string, c *color.Color) []string {
	colorizedStrings := make([]string, len(strs))
	for i, str := range strs {
		colorizedStrings[i] = c.Sprintf("%s", str)
	}
	return colorizedStrings
}

func convertToInterfaceSlice(patterns []string) []interface{} {
	args := []interface{}{}
	for _, pattern := range patterns {
		args = append(args, pattern)
	}
	return args
}

package marker

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// MatcherFunc returns a Match which contains information about found patterns
type MatcherFunc func(string) Match

// Match contains information about found patterns by MatcherFunc
type Match struct {
	Template string
	Patterns []string
}

// MatchAll creates a MatcherFunc that matches all patterns in given string
func MatchAll(pattern string) MatcherFunc {
	return func(str string) Match {
		count := strings.Count(str, pattern)
		return Match{
			Template: strings.ReplaceAll(str, pattern, "%s"),
			Patterns: fillSlice(make([]string, count), pattern),
		}
	}
}

// MatchN creates a MatcherFunc that matches first n patterns in given string
func MatchN(pattern string, n int) MatcherFunc {
	return func(str string) Match {
		count := min(n, strings.Count(str, pattern))
		return Match{
			Template: strings.Replace(str, pattern, "%s", n),
			Patterns: fillSlice(make([]string, count), pattern),
		}
	}
}

// MatchMultiple creates a MatcherFunc that matches all string patterns from given slice in given string
func MatchMultiple(patternsToMatch []string) MatcherFunc {
	return func(str string) Match {
		patternMatchIndexes := findPatternMatchIndexes(str, patternsToMatch)
		patterns := getPatternsInOrder(patternMatchIndexes)
		return Match{
			Template: replaceMultiple(str, patternsToMatch, "%s"),
			Patterns: patterns,
		}
	}
}

// MatchRegexp creates a MatcherFunc that matches given regexp in given string
func MatchRegexp(r *regexp.Regexp) MatcherFunc {
	return func(str string) Match {
		return Match{
			Template: r.ReplaceAllString(str, "%s"),
			Patterns: r.FindAllString(str, -1),
		}
	}
}

// MatchTimestamp creates a MatcherFunc that matches given time layout pattern in given string
func MatchTimestamp(layout string) MatcherFunc {
	return func(str string) Match {
		r := timestampLayoutRegexps[layout]
		return MatchRegexp(r)(str)
	}
}

// MatchSurrounded creates a MatcherFunc that matches the patterns surrounded by given opening and closure strings
func MatchSurrounded(opening string, closure string) MatcherFunc {
	return func(str string) Match {
		metaEscapedOpening := regexp.QuoteMeta(opening)
		metaEscapedClosure := regexp.QuoteMeta(closure)
		matchPattern := fmt.Sprintf("%s[^%s]*%s", metaEscapedOpening, metaEscapedOpening, metaEscapedClosure)
		r, _ := regexp.Compile(matchPattern)
		return MatchRegexp(r)(str)
	}
}

// MatchBracketSurrounded is a helper utility for easy matching of bracket surrounded text
func MatchBracketSurrounded() MatcherFunc {
	return MatchSurrounded("[", "]")
}

// MatchParensSurrounded is a helper utility for easy matching text surrounded in parentheses
func MatchParensSurrounded() MatcherFunc {
	return MatchSurrounded("(", ")")
}

// MatchEmail creates a MatcherFunc that matches emails which meets the conditions of RFC5322 standard
func MatchEmail() MatcherFunc {
	return func(str string) Match {
		return MatchRegexp(EmailRegexp)(str)
	}
}

var daysOfWeek = [14]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
	"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// MatchDaysOfWeek creates a MatcherFunc that matches days of the week in given string
func MatchDaysOfWeek() MatcherFunc {
	return func(str string) Match {
		return MatchMultiple(daysOfWeek[:])(str)
	}
}

func findPatternMatchIndexes(str string, patternsToMatch []string) map[int]string {
	patternMatchIndexes := make(map[int]string)
	for _, patternToMatch := range patternsToMatch {
		for strings.Contains(str, patternToMatch) {
			matchIndex := strings.Index(str, patternToMatch)
			str = strings.Replace(str, patternToMatch, "", 1)
			patternMatchIndexes[matchIndex] = patternToMatch
		}
	}
	return patternMatchIndexes
}

func getPatternsInOrder(patternMatchIndexes map[int]string) []string {
	matchIndexes := getKeys(patternMatchIndexes)
	sort.Ints(matchIndexes)
	patterns := make([]string, 0, len(patternMatchIndexes))
	for _, index := range matchIndexes {
		patterns = append(patterns, patternMatchIndexes[index])
	}
	return patterns
}

func getKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func replaceMultiple(str string, patterns []string, with string) string {
	for _, patternToMatch := range patterns {
		str = strings.ReplaceAll(str, patternToMatch, with)
	}
	return str
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fillSlice(s []string, v string) []string {
	for i := range s {
		s[i] = v
	}
	return s
}

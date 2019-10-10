package marker

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

var timestampLayoutRegexps = map[string]*regexp.Regexp{
	time.ANSIC:       ANSICRegexp,
	time.UnixDate:    UnixDateRegexp,
	time.RubyDate:    RubyDateRegexp,
	time.RFC822:      RFC822Regexp,
	time.RFC822Z:     RFC822ZRegexp,
	time.RFC850:      RFC850Regexp,
	time.RFC1123:     RFC1123Regexp,
	time.RFC1123Z:    RFC1123ZRegexp,
	time.RFC3339:     RFC3339Regexp,
	time.RFC3339Nano: RFC3339NanoRegexp,
	time.Kitchen:     KitchenRegexp,
	time.Stamp:       StampRegexp,
	time.StampMilli:  StampMilliRegexp,
	time.StampMicro:  StampMicroRegexp,
	time.StampNano:   StampNanoRegexp,
}

type MatcherFunc func(string) Match

// Match contains information about found patterns by MatcherFunc
type Match struct {
	Template string
	Patterns []string
}

// MatchAll returns a MatcherFunc that matches all patterns in given string
func MatchAll(pattern string) MatcherFunc {
	return func(str string) Match {
		count := strings.Count(str, pattern)
		return Match{
			Template: strings.ReplaceAll(str, pattern, "%s"),
			Patterns: fillSlice(make([]string, count), pattern),
		}
	}
}

// MatchN returns a MatcherFunc that matches first n patterns in given string
func MatchN(pattern string, n int) MatcherFunc {
	return func(str string) Match {
		count := min(n, strings.Count(str, pattern))
		return Match{
			Template: strings.Replace(str, pattern, "%s", n),
			Patterns: fillSlice(make([]string, count), pattern),
		}
	}
}

// MatchRegexp returns a MatcherFunc that matches regexp in given string
func MatchRegexp(r *regexp.Regexp) MatcherFunc {
	return func(str string) Match {
		return Match{
			Template: r.ReplaceAllString(str, "%s"),
			Patterns: r.FindAllString(str, -1),
		}
	}
}

// MatchTimestamp returns a MatcherFunc that matches a time layout pattern in given string
func MatchTimestamp(layout string) MatcherFunc {
	return func(str string) Match {
		r := timestampLayoutRegexps[layout]
		return MatchRegexp(r)(str)
	}
}

// MatchSurrounded takes in characters surrounding a given expected match and returns the match findings
func MatchSurrounded(charOne string, charTwo string) MatcherFunc {
	return func(str string) Match {
		quoteCharOne := regexp.QuoteMeta(charOne)
		quoteCharTwo := regexp.QuoteMeta(charTwo)
		matchPattern := fmt.Sprintf("%s[^%s]*%s", quoteCharOne, quoteCharOne, quoteCharTwo)
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

var daysOfWeek = [14]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
	"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// MatchDaysOfWeek returns a MatcherFunc that matches days of the week in given string
func MatchDaysOfWeek() MatcherFunc {
	return func(str string) Match {
		patternMatchIndexes := make(map[int]string)
		for _, day := range daysOfWeek {
			for strings.Contains(str, day) {
				matchIndex := strings.Index(str, day)
				str = strings.Replace(str, day, "%s", 1)
				patternMatchIndexes[matchIndex] = day
			}
		}
		matchIndexes := make([]int, 0, len(patternMatchIndexes))
		for matchKey := range patternMatchIndexes {
			matchIndexes = append(matchIndexes, matchKey)
		}
		sort.Ints(matchIndexes)
		pattern := make([]string, 0, len(patternMatchIndexes))
		for _, index := range matchIndexes {
			pattern = append(pattern, patternMatchIndexes[index])
		}
		return Match{
			Template: str,
			Patterns: pattern,
		}
	}
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

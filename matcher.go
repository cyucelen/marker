package marker

import (
	"regexp"
	"sort"
	"strings"
)

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

// MatchDaysOfWeek returns a MatcherFunc that matches days of the week in given string
func MatchDaysOfWeek() MatcherFunc {
	return func(str string) Match {
		daysOfWeek := [14]string{"monday", "Monday", "tuesday", "Tuesday", "wednesday", "Wednesday", "thursday", "Thursday", "friday", "Friday", "saturday", "Saturday", "sunday", "Sunday"}
		patternMatchIndexes := make(map[int]string)
		for  _, day := range daysOfWeek {
			for strings.Contains(str, day) {
				matchIndex := strings.Index(str, day)
				str = strings.Replace(str, day, "%s",1)
				patternMatchIndexes[matchIndex] = day
			}
		}
 	matchIndexes := make([]int, 0, len(patternMatchIndexes))
    	for matchKey,_ := range patternMatchIndexes {
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

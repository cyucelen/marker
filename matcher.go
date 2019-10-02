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

// MatchDaysofWeek returns a MatcherFunc that matches days of the week in given string
func MatchDaysofWeek() MatcherFunc {
	return func(str string) Match {
		var days [14]string = [14]string {"monday", "Monday", "tuesday", "Tuesday", "wednesday", "Wednesday", "thursday", "Thursday", "friday", "Friday", "saturday", "Saturday", "sunday", "Sunday"}
		pattern := make(map[int]string)
		newString := str
		for _, v := range days {
			if strings.Contains(newString, v) {
				count := strings.Count(newString, v)
				for i:=1; i<=count; i++ {
					loc := strings.Index(newString, v)
					newString = strings.Replace(newString, v, "%s",1)
					pattern[loc] = v
				}
			}
		}
		var keys []int
    		for k,_ := range pattern {
		        keys = append(keys, k)
		}
    		sort.Ints(keys)
		var pat []string
		for _, k := range keys{
			pat = append(pat, pattern[k])
		}
		return Match{
			Template: newString,
			Patterns: pat,
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

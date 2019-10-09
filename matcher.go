package marker

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	weekdaysAbv   = "((Mon)|(Tue)|(Wed)|(Thu)|(Fri)|(Sat)|(Sun))"
	weekdays      = "((Monday)|(Tuesday)|(Wednesday)|(Thursday)|(Friday)|(Saturday)|(Sunday))"
	monthsAbv     = "((Jan)|(Feb)|(Mar)|(Apr)|(May)|(Jun)|(Jul)|(Aug)|(Sep)|(Oct)|(Nov)|(Dec))"
	numericmonths = "(0[1-9]|1[0-2])"
	days          = "((_[1-9])|([1-2][0-9])|3[01])"               // _1 - 31
	daysWithZero  = "((0[1-9])|([1-2][0-9])|3[01])"               // 01 - 31
	hhmmss        = "(([0-1][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])" // 22:15:02
	hhmm          = "(([0-1][0-9]|2[0-3]):[0-5][0-9])"            // 22:15
	year          = "[0-9]{4}"                                    // 2006
	timezone      = "[A-Z]{3}"                                    // MST
	numericZone   = "-[0-9]{4}"                                   // -0300
	milli         = ".[0-9]{3}"
	micro         = ".[0-9]{6}"
	nano          = ".[0-9]{9}"
)

var (
	ANSICRegex 		 = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss,
		year))
	UnixDateRegex 	 = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss,
		timezone, year))
	RubyDateRegex 	 = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, daysWithZero,
		hhmmss, numericZone, year))
	RFC822Regex 	 = regexp.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm,
		timezone))
	RFC822ZRegex 	 = regexp.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm,
		numericZone))
	RFC850Regex 	 = regexp.MustCompile(fmt.Sprintf("%s, %s-%s-[0-9]{2} %s %s", weekdays, daysWithZero,
		monthsAbv, hhmmss, timezone))
	RFC1123Regex 	 = regexp.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv,
		year, hhmmss, timezone))
	RFC1123ZRegex 	 = regexp.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv,
		year, hhmmss, numericZone))
	RFC3339Regex 	 = regexp.MustCompile(fmt.Sprintf("%s-%s-%sT%sZ%s", year, numericmonths, daysWithZero,
		hhmmss, hhmm))
	RFC3339NanoRegex = regexp.MustCompile(fmt.Sprintf("%s-%s-%sT%s%sZ%s", year, numericmonths, daysWithZero,
		hhmmss, nano, hhmm))
	KitchenRegex 	 = regexp.MustCompile("(([0-1]?[0-9]|2[0-3]):[0-5][0-9][P|A]M)")
	StampRegex 		 = regexp.MustCompile(fmt.Sprintf("%s %s %s", monthsAbv, days, hhmmss))
	StampMilliRegex  = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, milli))
	StampMicroRegex  = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, micro))
	StampNanoRegex   = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, nano))
)

var layoutRegexps = map[string]*regexp.Regexp{
	time.ANSIC: 	  ANSICRegex,
	time.UnixDate: 	  UnixDateRegex,
	time.RubyDate: 	  RubyDateRegex,
	time.RFC822: 	  RFC822Regex,
	time.RFC822Z: 	  RFC822ZRegex,
	time.RFC850: 	  RFC850Regex,
	time.RFC1123: 	  RFC1123Regex,
	time.RFC1123Z: 	  RFC1123ZRegex,
	time.RFC3339: 	  RFC3339Regex,
	time.RFC3339Nano: RFC3339NanoRegex,
	time.Kitchen: 	  KitchenRegex,
	time.Stamp: 	  StampRegex,
	time.StampMilli:  StampMilliRegex,
	time.StampMicro:  StampMicroRegex,
	time.StampNano:   StampNanoRegex,
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
func MatchTimestamp(layout string) MatcherFunc{
	return func(str string) Match {
		r := layoutRegexps[layout]

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

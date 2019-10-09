package marker

import (
	"fmt"
	"regexp"
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
	ANSICRegexp       = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss, year))
	UnixDateRegexp    = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, days, hhmmss, timezone, year))
	RubyDateRegexp    = regexp.MustCompile(fmt.Sprintf("%s %s %s %s %s %s", weekdaysAbv, monthsAbv, daysWithZero, hhmmss, numericZone, year))
	RFC822Regexp      = regexp.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm, timezone))
	RFC822ZRegexp     = regexp.MustCompile(fmt.Sprintf("%s %s [0-9]{2} %s %s", daysWithZero, monthsAbv, hhmm, numericZone))
	RFC850Regexp      = regexp.MustCompile(fmt.Sprintf("%s, %s-%s-[0-9]{2} %s %s", weekdays, daysWithZero, monthsAbv, hhmmss, timezone))
	RFC1123Regexp     = regexp.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv, year, hhmmss, timezone))
	RFC1123ZRegexp    = regexp.MustCompile(fmt.Sprintf("%s, %s %s %s %s %s", weekdaysAbv, daysWithZero, monthsAbv, year, hhmmss, numericZone))
	RFC3339Regexp     = regexp.MustCompile(fmt.Sprintf("%s-%s-%sT%sZ%s", year, numericmonths, daysWithZero, hhmmss, hhmm))
	RFC3339NanoRegexp = regexp.MustCompile(fmt.Sprintf("%s-%s-%sT%s%sZ%s", year, numericmonths, daysWithZero, hhmmss, nano, hhmm))
	KitchenRegexp     = regexp.MustCompile("(([0-1]?[0-9]|2[0-3]):[0-5][0-9][P|A]M)")
	StampRegexp       = regexp.MustCompile(fmt.Sprintf("%s %s %s", monthsAbv, days, hhmmss))
	StampMilliRegexp  = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, milli))
	StampMicroRegexp  = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, micro))
	StampNanoRegexp   = regexp.MustCompile(fmt.Sprintf("%s %s %s%s", monthsAbv, days, hhmmss, nano))
)

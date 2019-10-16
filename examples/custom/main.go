package main

import (
	"fmt"
	"strings"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func MatchFirst(pattern string) marker.MatcherFunc {
	return func(str string) marker.Match {
		return marker.Match{
			// replace first matching pattern with %s
			Template: strings.Replace(str, pattern, "%s", 1),
			// patterns to be colorized by Mark, in order
			Patterns: []string{pattern},
		}
	}
}

func main() {
	boringLog := "[INFO] Nobody wants to read pale [INFO] tags."
	brilliantLog := marker.Mark(boringLog, MatchFirst("[INFO]"), color.New(color.FgBlue))
	fmt.Println(brilliantLog)
}

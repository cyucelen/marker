package main

import (
	"fmt"
	"regexp"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func main() {
	rhyme := "I scream, you all scream, we all scream for ice cream."
	r, _ := regexp.Compile("([a-z]?cream)")
	careAboutCream := marker.Mark(rhyme, marker.MatchRegexp(r), color.New(color.FgYellow))
	fmt.Println(careAboutCream)
}

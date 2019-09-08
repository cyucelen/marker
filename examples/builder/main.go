package main

import (
	"fmt"
	"regexp"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func main() {
	rhyme := "I scream, you all scream, we all scream for ice cream."
	b := &marker.MarkBuilder{}
	r, _ := regexp.Compile("([a-z]?cream)")

	markedWithBuilder := b.SetString(rhyme).
		Mark(marker.MatchN("for ice", 1), color.New(color.FgRed)).
		Mark(marker.MatchAll("all"), color.New(color.FgMagenta)).
		Mark(marker.MatchRegexp(r), color.New(color.FgYellow)).
		Build()

	fmt.Println(markedWithBuilder)
}

package main

import (
	"fmt"
	"regexp"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var magentaFg = color.New(color.FgMagenta)
var hiMagentaFg = color.New(color.FgHiMagenta)
var redFg = color.New(color.FgRed)
var blueFg = color.New(color.FgBlue)
var blackFg = color.New(color.FgBlack)
var greenFg = color.New(color.FgGreen)
var yellowFg = color.New(color.FgYellow)
var whiteFg = color.New(color.FgWhite)

func main() {
	rhyme := "I scream, you all scream, we all scream for ice cream."

	rhymeMarkExample := "Mark All \"all\":\t\t\t\t\t\t" + rhyme
	allMarked := marker.Mark(rhymeMarkExample, marker.MatchAll("all"), magentaFg)
	fmt.Println(allMarked)

	rhymeMarkExample = "Mark All \"all\" and \"ice\":\t\t\t\t" + rhyme
	allMarked = marker.Mark(rhymeMarkExample, marker.MatchAll("all"), greenFg)
	allIceMarked := marker.Mark(allMarked, marker.MatchAll("ice"), hiMagentaFg.Add(color.BgWhite))
	fmt.Println(allIceMarked)

	r, _ := regexp.Compile("([a-z]?cream)")
	markedWithRegexp := marker.Mark(rhyme, marker.MatchRegexp(r), whiteFg.Add(color.BgHiBlue))
	regexpExampleHeader := fmt.Sprintf("Mark Regexp \"%s\":\t\t\t", whiteFg.Add(color.BgHiBlue).Sprint("([a-z]?cream)"))
	fmt.Println(regexpExampleHeader + markedWithRegexp)

	b := &marker.MarkBuilder{}
	markedWithBuilder := b.SetString(rhyme).
		Mark(marker.MatchN("for ice", 1), redFg).
		Mark(marker.MatchAll("all"), magentaFg).
		Mark(marker.MatchRegexp(r), blueFg).
		Build()
	builderExampleHeader := fmt.Sprintf("Mark \"%s\", \"%s\", \"%s\" :\t",
		color.New(color.FgRed).Sprint("for ice"), magentaFg.Sprint("all"), blueFg.Sprint("([a-z]?cream)"))
	fmt.Println(builderExampleHeader + markedWithBuilder)
}

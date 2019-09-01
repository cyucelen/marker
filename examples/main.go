package main

import (
	"fmt"
	"github.com/cyucelen/marker"
	"github.com/fatih/color"
	"regexp"
)

var magentaFg = color.New(color.FgMagenta)
var hiMagentaFg = color.New(color.FgHiMagenta)
var redFg = color.New(color.FgRed)
var blueFg = color.New(color.FgBlue)
var blackFg = color.New(color.FgBlack)
var greenFg = color.New(color.FgGreen)

func main() {
	rhyme := "I scream, you all scream, we all scream for ice cream."

	rhymeMarkExample := "Mark All \"all\":\t\t\t\t\t\t" + rhyme
	allMarked := marker.Mark(rhymeMarkExample, marker.MatchAll("all"), magentaFg)
	fmt.Println(allMarked)

	rhymeMarkExample = "Mark All \"all\" and \"ice\":\t\t\t\t" + rhyme
	allMarked = marker.Mark(rhymeMarkExample, marker.MatchAll("all"), magentaFg)
	allIceMarked := marker.Mark(allMarked, marker.MatchAll("ice"), hiMagentaFg.Add(color.BgBlack))
	fmt.Println(allIceMarked)

	r, _ := regexp.Compile("([a-z]?cream)")
	markedWithRegexp := marker.Mark(rhyme, marker.MatchRegexp(r), blackFg.Add(color.BgGreen))
	regexpExampleHeader := fmt.Sprintf("Mark Regexp \"%s\":\t\t\t\t", blackFg.Add(color.BgGreen).Sprint("([a-z]?cream)"))
	fmt.Println(regexpExampleHeader + markedWithRegexp)

	b := &marker.MarkBuilder{}
	markedWithBuilder := b.SetString(rhyme).
		Mark(marker.MatchN("for ice", 1), redFg).
		Mark(marker.MatchAll("all"), magentaFg).
		Mark(marker.MatchRegexp(r), blackFg.Add(color.BgHiMagenta)).
		Build()
	builderExampleHeader := fmt.Sprintf("Mark \"%s\", \"%s\", \"%s\" (w/ builder):\t",
		redFg.Sprint("for ice"), magentaFg.Sprint("all"), blackFg.Add(color.BgHiMagenta).Sprint("([a-z]?cream)"))
	fmt.Println(builderExampleHeader + markedWithBuilder)
}

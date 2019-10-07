package main

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var magentaFg = color.New(color.FgMagenta)
var redFg = color.New(color.FgRed)
var blueFg = color.New(color.FgBlue)

func main() {
	sentence := "I pull out things surrounded by abcWHOA COLORSdef"
	markedSurrounded := marker.Mark(sentence, marker.MatchSurrounded("abc", "def"), magentaFg)
	fmt.Println(markedSurrounded)

	sentence = "[INFO] This is what log lines look like"
	markedSurrounded = marker.Mark(sentence, marker.MatchBracketSurrounded(), redFg)
	fmt.Println(markedSurrounded)

	sentence = "[ERROR] This is what (parens) lines look like"
	markedSurrounded = marker.Mark(sentence, marker.MatchParensSurrounded(), blueFg)
	fmt.Println(markedSurrounded)
}

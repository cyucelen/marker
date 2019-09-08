package main

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func main() {
	aristotleQuote := "The more you know, the more you realize you don't know."
	emphasized := marker.Mark(aristotleQuote, marker.MatchAll("know"), color.New(color.FgRed))
	fmt.Println(emphasized)
}

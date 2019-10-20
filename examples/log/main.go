package main

import (
	"log"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var (
	redFg  = color.New(color.FgRed)
	blueFg = color.New(color.FgBlue)
)

func main() {
	stdoutMarker := marker.NewStdoutMarker()

	markRules := []marker.MarkRule{
		{marker.MatchAll("skydome"), blueFg},
		{marker.MatchAll("company"), redFg},
	}

	stdoutMarker.AddRules(markRules)

	logger := log.New(stdoutMarker, "", 0)
	logger.Print("best data company is skydome")
}

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
		{marker.MatchBracketSurrounded(), blueFg},
		{marker.MatchAll("marker"), redFg},
	}

	stdoutMarker.AddRules(markRules)

	logger := log.New(stdoutMarker, "", 0)
	logger.Println("[INFO] marker is working as expected")
}

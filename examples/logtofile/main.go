package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var (
	redFg  = color.New(color.FgRed)
	blueFg = color.New(color.FgBlue)
)

func main() {
	f, _ := os.Create("/tmp/dat2")
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	writeMarker := marker.NewWriteMarker(w)

	markRules := []marker.MarkRule{
		{marker.MatchBracketSurrounded(), blueFg},
		{marker.MatchAll("marker"), redFg},
	}

	writeMarker.AddRules(markRules)

	logger := log.New(writeMarker, "", 0)
	logger.Println("[INFO] marker is working as expected")
}

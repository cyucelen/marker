package main

import (
	"fmt"
	"time"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var blueFg = color.New(color.FgBlue)

func main() {
	goodOldTimes := "2006-01-02T15:04:05Z07:00 [INFO] Loading King of Fighters '97 ROM"
	timestampMarked := marker.Mark(goodOldTimes, marker.MatchTimestamp(time.RFC3339), blueFg)
	fmt.Println(timestampMarked)
}

package main

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func main() {
	message := "This is a message by dev@gmail.com and john@doe.io"
	brilliantLog := marker.Mark(message, marker.MatchEmail(), color.New(color.FgRed))
	fmt.Println(brilliantLog)
}

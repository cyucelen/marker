package main

import (
	"fmt"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func main() {
	boringLog := "[INFO] Nobody wants to read pale [INFO] tags."
	brilliantLog := marker.Mark(boringLog, marker.MatchN("[INFO]", 1), color.New(color.FgBlack))
	fmt.Println(brilliantLog)
}

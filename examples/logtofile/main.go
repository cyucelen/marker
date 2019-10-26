package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

var (
	magentaFg = color.New(color.FgMagenta)
	blueFg    = color.New(color.FgBlue)
)

func main() {
	f, _ := os.Create("/tmp/awesome.log")
	w := bufio.NewWriter(f)

	writeMarker := marker.NewWriteMarker(w)

	markRules := []marker.MarkRule{
		{marker.MatchBracketSurrounded(), blueFg},
		{marker.MatchAll("marker"), magentaFg},
	}

	writeMarker.AddRules(markRules)

	logger := log.New(writeMarker, "", 0)
	logger.Println("[INFO] colorful logs even in files, marker to mark them all!")

	w.Flush()
	f.Close()

	output := catFile("/tmp/awesome.log") // cat /tmp/dat2
	fmt.Print(output)
}

func catFile(file string) string {
	cmd := exec.Command("cat", file)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(cmdOutput.Bytes())
}

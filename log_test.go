package marker

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

type MockLogOut struct {
	actualLog string
}

func (m *MockLogOut) Write(p []byte) (n int, err error) {
	m.actualLog = string(p)
	return len(p), nil
}

func Test_New(t *testing.T) {
	stdoutMarker := NewStdoutMarker()
	assert.Equal(t, stdoutMarker.out, os.Stdout)

	mockLogOut := &MockLogOut{}
	stdoutMarker = NewStdoutMarker(SetOutput(mockLogOut))
	assert.Equal(t, stdoutMarker.out, mockLogOut)
}

func Test_Write(t *testing.T) {
	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()
	blueFg := color.New(color.FgBlue)
	blueFg.EnableColor()
	blue := blueFg.SprintFunc()

	mockOut := &MockLogOut{}
	stdoutMarker := NewStdoutMarker(SetOutput(mockOut))

	stdoutMarker.AddRule(MarkRule{MatchAll("skydome"), redFg}).AddRule(MarkRule{MatchAll("data"), redFg})

	logger := log.New(stdoutMarker, "", 0)
	logger.Print("best data company is skydome")

	expectedLog := fmt.Sprintf("best %s company is %s\n", red("data"), red("skydome"))
	assert.Equal(t, expectedLog, mockOut.actualLog)

	// Testing the mark order here since we cannot assert function equality (https://golang.org/ref/spec#Comparison_operators)
	newRules := []MarkRule{
		{MatchAll("skydome"), blueFg}, // blue should override red because of order
		{MatchAll("company"), redFg},
	}
	stdoutMarker.AddRules(newRules)

	expectedLog = fmt.Sprintf("best %s %s is %s\n", red("data"), red("company"), red(blue("skydome")))
	logger.Print("best data company is skydome")
	assert.Equal(t, expectedLog, mockOut.actualLog)

}

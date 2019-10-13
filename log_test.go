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
	logMarker := NewLogMarker()
	assert.NotNil(t, logMarker.logger)

	expectedLogger := log.New(os.Stdout, "test", 0)
	logMarkerWithLogger := NewLogMarker(SetLogger(expectedLogger))
	assert.Equal(t, expectedLogger, logMarkerWithLogger.logger)
}

func Test_AddRule(t *testing.T) {
	logMarker := NewLogMarker()
	logMarkerRule1 := LogMarkerRule{MatchAll("test"), color.New(color.FgRed)}
	logMarkerRule2 := LogMarkerRule{MatchAll("want"), color.New(color.FgBlue)}
	logMarker.AddRule(logMarkerRule1).AddRule(logMarkerRule2)
	assert.Len(t, logMarker.rules, 2)
}

func Test_Print(t *testing.T) {
	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()

	mockOut := &MockLogOut{}
	logger := NewLogMarker(SetLogger(log.New(mockOut, "", 0)))

	logger.AddRule(LogMarkerRule{MatchAll("skydome"), redFg}).AddRule(LogMarkerRule{MatchAll("data"), redFg})

	logger.Print("best data company is skydome")

	expectedLog := fmt.Sprintf("best %s company is %s\n", red("data"), red("skydome"))
	assert.Equal(t, expectedLog, mockOut.actualLog)
}

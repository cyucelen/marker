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
	logMarker := NewLogger()
	assert.NotNil(t, logMarker.logger)

	expectedLogger := log.New(os.Stdout, "test", 0)
	logMarkerWithLogger := NewLogger(SetLogger(expectedLogger))
	assert.Equal(t, expectedLogger, logMarkerWithLogger.logger)
}

func Test_AddRule(t *testing.T) {
	logger := NewLogger()
	markRule := MarkRule{MatchAll("test"), color.New(color.FgRed)}
	logger.AddRule(markRule)
	assert.Len(t, logger.rules, 1)

	markRule = MarkRule{MatchAll("want"), color.New(color.FgBlue)}
	logger.AddRule(markRule)
	assert.Len(t, logger.rules, 2)
}

func Test_AddRules(t *testing.T) {
	logger := NewLogger()
	markRules := []MarkRule{
		{MatchBracketSurrounded(), color.New(color.FgGreen)},
		{MatchParensSurrounded(), color.New(color.BgBlack)},
		{MatchEmail(), color.New(color.FgCyan)},
	}
	logger.AddRules(markRules)
	assert.Len(t, logger.rules, 3)

	logger.AddRules(markRules[:2])
	assert.Len(t, logger.rules, 5)
}

func Test_Print(t *testing.T) {
	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()

	mockOut := &MockLogOut{}
	logger := NewLogger(SetLogger(log.New(mockOut, "", 0)))

	logger.AddRule(MarkRule{MatchAll("skydome"), redFg}).AddRule(MarkRule{MatchAll("data"), redFg})

	logger.Print("best data company is skydome")

	expectedLog := fmt.Sprintf("best %s company is %s\n", red("data"), red("skydome"))
	assert.Equal(t, expectedLog, mockOut.actualLog)
}

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

func Test_AddRule(t *testing.T) {
	stdoutMarker := NewStdoutMarker()
	markRule := MarkRule{MatchAll("test"), color.New(color.FgRed)}
	stdoutMarker.AddRule(markRule)
	assertMarkRuleEqual(t, stdoutMarker.rules[0], markRule)

	markRule = MarkRule{MatchAll("want"), color.New(color.FgBlue)}
	stdoutMarker.AddRule(markRule)
	assertMarkRuleEqual(t, stdoutMarker.rules[1], markRule)
}

func Test_AddRules(t *testing.T) {
	stdoutMarker := NewStdoutMarker()
	expectedRules := []MarkRule{
		{MatchBracketSurrounded(), color.New(color.FgGreen)},
		{MatchParensSurrounded(), color.New(color.BgBlack)},
		{MatchEmail(), color.New(color.FgCyan)},
	}
	stdoutMarker.AddRules(expectedRules)
	assertMarkRuleSliceEqual(t, expectedRules, stdoutMarker.rules)

	newRules := []MarkRule{
		{MatchDaysOfWeek(), color.New(color.FgHiBlue)},
		{MatchSurrounded("[", ")"), color.New(color.FgHiRed)},
	}
	stdoutMarker.AddRules(newRules)
	expectedRules = append(expectedRules, newRules...)
	assertMarkRuleSliceEqual(t, expectedRules, stdoutMarker.rules)
}

func Test_Write(t *testing.T) {
	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()

	mockOut := &MockLogOut{}
	stdoutMarker := NewStdoutMarker(SetOutput(mockOut))

	stdoutMarker.AddRule(MarkRule{MatchAll("skydome"), redFg}).AddRule(MarkRule{MatchAll("data"), redFg})

	logger := log.New(stdoutMarker, "", 0)
	logger.Print("best data company is skydome")

	expectedLog := fmt.Sprintf("best %s company is %s\n", red("data"), red("skydome"))
	assert.Equal(t, expectedLog, mockOut.actualLog)
}

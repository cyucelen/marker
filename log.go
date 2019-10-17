package marker

import (
	"io"
	"os"

	"github.com/fatih/color"
)

type StdoutMarkerOption func(*StdoutMarker)

type MarkRule struct {
	Matcher MatcherFunc
	Color   *color.Color
}

type StdoutMarker struct {
	rules []MarkRule
	out   io.Writer
}

func NewStdoutMarker(options ...StdoutMarkerOption) *StdoutMarker {
	logMarker := &StdoutMarker{out: os.Stdout}
	for _, option := range options {
		option(logMarker)
	}
	return logMarker
}

func (s *StdoutMarker) AddRule(rule MarkRule) *StdoutMarker {
	s.rules = append(s.rules, rule)
	return s
}

func (s *StdoutMarker) AddRules(rules []MarkRule) {
	s.rules = append(s.rules, rules...)
}

func (s StdoutMarker) Write(p []byte) (n int, err error) {
	marked := string(p)
	for _, rule := range s.rules {
		marked = Mark(marked, rule.Matcher, rule.Color)
	}
	return s.out.Write([]byte(marked))
}

func SetOutput(out io.Writer) StdoutMarkerOption {
	return func(stdoutWriter *StdoutMarker) {
		stdoutWriter.out = out
	}
}

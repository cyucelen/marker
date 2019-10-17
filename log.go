package marker

import (
	"io"
	"os"

	"github.com/fatih/color"
)

// StdoutMarkerOption is functional option type for StdoutMarker
type StdoutMarkerOption func(*StdoutMarker)

// MarkRule contains marking information to be applied on log stream
type MarkRule struct {
	Matcher MatcherFunc
	Color   *color.Color
}

// StdoutMarker contains specified rules for applying them on output
type StdoutMarker struct {
	rules []MarkRule
	out   io.Writer
}

// NewStdoutMarker creates a StdoutMarker with default out as os.Stdout
func NewStdoutMarker(options ...StdoutMarkerOption) *StdoutMarker {
	logMarker := &StdoutMarker{out: os.Stdout}
	for _, option := range options {
		option(logMarker)
	}
	return logMarker
}

// AddRule appends a rule to StdoutMarker and returns itself
func (s *StdoutMarker) AddRule(rule MarkRule) *StdoutMarker {
	s.rules = append(s.rules, rule)
	return s
}

// AddRules appends all rules from given slice to StdoutMarker rules
func (s *StdoutMarker) AddRules(rules []MarkRule) {
	s.rules = append(s.rules, rules...)
}

// Write marks the text with specified rules and writes the output to specifed out
func (s StdoutMarker) Write(p []byte) (n int, err error) {
	marked := string(p)
	for _, rule := range s.rules {
		marked = Mark(marked, rule.Matcher, rule.Color)
	}
	return s.out.Write([]byte(marked))
}

// SetOutput is a Functional Option for specifying output of StdoutMarker
func SetOutput(out io.Writer) StdoutMarkerOption {
	return func(stdoutWriter *StdoutMarker) {
		stdoutWriter.out = out
	}
}

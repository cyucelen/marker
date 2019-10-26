package marker

import (
	"io"
	"os"

	"github.com/fatih/color"
)

// WriteMarkerOption is functional option type for WriteMarker
type WriteMarkerOption func(*WriteMarker)

// MarkRule contains marking information to be applied on log stream
type MarkRule struct {
	Matcher MatcherFunc
	Color   *color.Color
}

// WriteMarker contains specified rules for applying them on output
type WriteMarker struct {
	rules []MarkRule
	out   io.Writer
}

// NewWriteMarker creates a Marker that writes out to the given io.Writer
func NewWriteMarker(writer io.Writer) *WriteMarker {
	logMarker := &WriteMarker{out: writer}
	return logMarker
}

// NewStdoutMarker creates a WriteMarker with default out as os.Stdout
func NewStdoutMarker() *WriteMarker {
	return NewWriteMarker(os.Stdout)
}

// AddRule appends a rule to WriteMarker and returns itself
func (s *WriteMarker) AddRule(rule MarkRule) *WriteMarker {
	s.rules = append(s.rules, rule)
	return s
}

// AddRules appends all rules from given slice to WriteMarker rules
func (s *WriteMarker) AddRules(rules []MarkRule) {
	s.rules = append(s.rules, rules...)
}

// Write marks the text with specified rules and writes the output to specifed out
func (s WriteMarker) Write(p []byte) (n int, err error) {
	marked := string(p)
	for _, rule := range s.rules {
		marked = Mark(marked, rule.Matcher, rule.Color)
	}
	return s.out.Write([]byte(marked))
}

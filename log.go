package marker

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type LoggerOption func(*Logger)

type MarkRule struct {
	Matcher MatcherFunc
	Color   *color.Color
}

type Logger struct {
	rules  []MarkRule
	logger *log.Logger
}

func NewLogger(options ...LoggerOption) *Logger {
	logMarker := &Logger{logger: log.New(os.Stdout, "", 0)}
	for _, option := range options {
		option(logMarker)
	}
	return logMarker
}

func SetLogger(logger *log.Logger) LoggerOption {
	return func(logMarker *Logger) {
		logMarker.logger = logger
	}
}

func (l *Logger) AddRule(rule MarkRule) *Logger {
	l.rules = append(l.rules, rule)
	return l
}

func (l *Logger) Print(v ...interface{}) {
	marked := fmt.Sprint(v...)
	for _, rule := range l.rules {
		marked = Mark(marked, rule.Matcher, rule.Color)
	}
	l.logger.Print(marked)
}

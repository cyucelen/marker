package marker

import (
	"github.com/fatih/color"
)

type MarkBuilder struct {
	str string
}

func (m *MarkBuilder) SetString(str string) *MarkBuilder {
	m.str = str
	return m
}

func (m *MarkBuilder) Mark(matcherFunc MatcherFunc, c *color.Color) *MarkBuilder {
	m.str = Mark(m.str, matcherFunc, c)
	return m
}

func (m *MarkBuilder) Build() string {
	return m.str
}

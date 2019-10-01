package marker

import (
	"github.com/fatih/color"
)

MarkBuilder is a better and neater way to mark different patterns of the string
type MarkBuilder struct {
	str string
}

SetString sets the first parameter as the string that is going to be marked
func (m *MarkBuilder) SetString(str string) *MarkBuilder {
	m.str = str
	return m
}

Mark marks the string with the matcher function and color
func (m *MarkBuilder) Mark(matcherFunc MatcherFunc, c *color.Color) *MarkBuilder {
	m.str = Mark(m.str, matcherFunc, c)
	return m
}

Build returns the marked string
func (m *MarkBuilder) Build() string {
	return m.str
}

package marker

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Mark(t *testing.T) {
	blueFg := color.New(color.FgBlue)
	blueFg.EnableColor()
	blue := blueFg.SprintFunc()

	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()

	tests := []struct {
		text     string
		matcher  MatcherFunc
		expected string
		color    *color.Color
	}{
		{
			text:  "Skydome is a data company.",
			color: blueFg,
			matcher: func(str string) Match {
				return Match{Template: "%s is a data company.", Patterns: []string{"Skydome"}}
			},
			expected: fmt.Sprintf("%s is a data company.", blue("Skydome")),
		},
		{
			text:  "Skydome is Skydome. Give yourself freedom.",
			color: redFg,
			matcher: func(str string) Match {
				return Match{Template: "%s is %s. Give yourself freedom.", Patterns: []string{"Skydome", "Skydome"}}
			},
			expected: fmt.Sprintf("%s is %s. Give yourself freedom.", red("Skydome"), red("Skydome")),
		},
	}

	for _, testCase := range tests {
		actual := Mark(testCase.text, testCase.matcher, testCase.color)
		assert.Equal(t, testCase.expected, actual)
	}
}

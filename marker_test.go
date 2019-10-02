package marker

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
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

func Benchmark_Mark(b *testing.B) {
	blueFg := color.New(color.FgBlue)
	blueFg.EnableColor()
	b.ReportAllocs()

	data := struct {
		text     string
		matcher  MatcherFunc
		expected string
		color    *color.Color
	}{
		text:  "Skydome is a data company.",
		color: blueFg,
		matcher: func(str string) Match {
			return Match{Template: "%s is a data company.", Patterns: []string{"Skydome"}}
		},
	}

	for i := 0; i < b.N; i++ {
		Mark(data.text, data.matcher, data.color)
	}
}

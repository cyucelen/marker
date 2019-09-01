package marker

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func Test_Builder(t *testing.T) {
	blueFg := color.New(color.FgBlue)
	blueFg.EnableColor()
	blue := blueFg.SprintFunc()

	redFg := color.New(color.FgRed)
	redFg.EnableColor()
	red := redFg.SprintFunc()

	b := MarkBuilder{}

	actualString := b.SetString("Skydome is a data company.").
		Mark(MatchAll("Skydome"), blueFg).
		Mark(MatchAll("data"), redFg).
		Build()

	expectedString := fmt.Sprintf("%s is a %s company.", blue("Skydome"), red("data"))

	assert.Equal(t, expectedString, actualString)
}

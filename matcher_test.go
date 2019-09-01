package marker

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_MatchAll(t *testing.T) {
	str := "Skydome is Skydome"
	actualMatch := MatchAll("Skydome")(str)
	expectedMatch := Match{Template: "%s is %s", Patterns: []string{"Skydome", "Skydome"}}

	assert.Equal(t, actualMatch, expectedMatch)
}

func Test_MatchN(t *testing.T) {
	str := "Skydome is Skydome"
	actualMatch := MatchN("Skydome", 1)(str)
	expectedMatch := Match{Template: "%s is Skydome", Patterns: []string{"Skydome"}}
	assert.Equal(t, expectedMatch, actualMatch)

	actualMatch = MatchN("Skydome", 2)(str)
	expectedMatch = Match{Template: "%s is %s", Patterns: []string{"Skydome", "Skydome"}}
	assert.Equal(t, expectedMatch, actualMatch)

	actualMatch = MatchN("Skydome", 3)(str)
	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_MatchRegexp(t *testing.T) {
	str := "I scream, you all scream, we all scream for ice cream."

	r, _ := regexp.Compile("([a-z]?cream)")
	actualMatch := MatchRegexp(r)(str)
	expectedMatch := Match{Template: "I %s, you all %s, we all %s for ice %s.", Patterns: []string{"scream", "scream", "scream", "cream"}}

	assert.Equal(t, actualMatch, expectedMatch)
}

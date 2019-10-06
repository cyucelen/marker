package marker

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MatchAll(t *testing.T) {
	str := "Skydome is Skydome"
	actualMatch := MatchAll("Skydome")(str)
	expectedMatch := Match{Template: "%s is %s", Patterns: []string{"Skydome", "Skydome"}}

	assert.Equal(t, expectedMatch, actualMatch)
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

	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_MatchSurrounded(t *testing.T) {
	str := "[ERROR] This is a -debug- message -(and it's okay)- [INFO] --test--"

	actualMatch := MatchSurrounded("-", "-")(str)

	expectedMatch := Match{
		Template: "[ERROR] This is a %s message %s [INFO] %stest%s",
		Patterns: []string{"-debug-", "-(and it's okay)-", "--", "--"},
	}

	assert.Equal(t, expectedMatch, actualMatch)

	str = "abcMULTIPLE CHARACTERSdef whoa"

	actualMatch = MatchSurrounded("abc", "def")(str)

	expectedMatch = Match{
		Template: "%s whoa",
		Patterns: []string{"abcMULTIPLE CHARACTERSdef"},
	}

	assert.Equal(t, expectedMatch, actualMatch)

	str = "[[DOUBLE CHARACTERS]]"

	actualMatch = MatchSurrounded("[[", "]]")(str)

	expectedMatch = Match{
		Template: "%s",
		Patterns: []string{"[[DOUBLE CHARACTERS]]"},
	}

	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_MatchBracketSurrounded(t *testing.T) {
	str := "[ERROR] This is a -debug- message (and it's okay) [INFO] --test--"

	actualMatch := MatchBracketSurrounded()(str)

	expectedMatch := Match{
		Template: "%s This is a -debug- message (and it's okay) %s --test--",
		Patterns: []string{"[ERROR]", "[INFO]"},
	}

	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_MatchParensSurrounded(t *testing.T) {
	str := "[ERROR] This is a -debug- message (and it's okay) [INFO] --test--"

	actualMatch := MatchParensSurrounded()(str)

	expectedMatch := Match{
		Template: "[ERROR] This is a -debug- message %s [INFO] --test--",
		Patterns: []string{"(and it's okay)"},
	}

	assert.Equal(t, expectedMatch, actualMatch)
}

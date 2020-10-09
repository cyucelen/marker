package marker

import (
	"regexp"
	"testing"
	"time"

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

func Test_MatchTimestamp(t *testing.T) {
	t.Parallel()

	t.Run("ANSIC", func(t *testing.T) {
		str := "Current timestamp is Mon Jan 31 20:59:00 2006"
		match := MatchTimestamp(time.ANSIC)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Mon Jan 31 20:59:00 2006"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("UnixDate", func(t *testing.T) {
		str := "Current timestamp is Mon Jan _2 03:04:05 MST 2006"
		match := MatchTimestamp(time.UnixDate)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Mon Jan _2 03:04:05 MST 2006"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RubyDate", func(t *testing.T) {
		str := "Current timestamp is Fri Feb 04 03:04:05 -0300 2006"
		match := MatchTimestamp(time.RubyDate)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Fri Feb 04 03:04:05 -0300 2006"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC822", func(t *testing.T) {
		str := "Current timestamp is 02 Jan 06 05:04 MST"
		match := MatchTimestamp(time.RFC822)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"02 Jan 06 05:04 MST"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC822Z", func(t *testing.T) {
		str := "Current timestamp is 02 Jan 06 15:04 -0300"
		match := MatchTimestamp(time.RFC822Z)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"02 Jan 06 15:04 -0300"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC850", func(t *testing.T) {
		str := "Current timestamp is Saturday, 07-Aug-06 22:04:59 MST"
		match := MatchTimestamp(time.RFC850)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Saturday, 07-Aug-06 22:04:59 MST"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC1123", func(t *testing.T) {
		str := "Current timestamp is Mon, 02 Jan 2006 15:04:05 MST"
		match := MatchTimestamp(time.RFC1123)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Mon, 02 Jan 2006 15:04:05 MST"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC1123Z", func(t *testing.T) {
		str := "Current timestamp is Mon, 02 Jan 2006 15:04:05 -0300"
		match := MatchTimestamp(time.RFC1123Z)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Mon, 02 Jan 2006 15:04:05 -0300"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC3339", func(t *testing.T) {
		str := "Current timestamp is 2006-01-02T15:04:05Z07:00"
		match := MatchTimestamp(time.RFC3339)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"2006-01-02T15:04:05Z07:00"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("RFC3339Nano", func(t *testing.T) {
		str := "Current timestamp is 2006-01-02T15:04:05.999999999Z07:00"
		match := MatchTimestamp(time.RFC3339Nano)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"2006-01-02T15:04:05.999999999Z07:00"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("Kitchen", func(t *testing.T) {
		str := "Current timestamp is 2:15PM"
		match := MatchTimestamp(time.Kitchen)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"2:15PM"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("Stamp", func(t *testing.T) {
		str := "Current timestamp is Jan _2 15:04:05 and Jan _2 15:04:10"
		match := MatchTimestamp(time.Stamp)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %sand %s",
			Patterns: []string{"Jan _2 15:04:05 ", "Jan _2 15:04:10"},
		}

		assert.Equal(t, expectedMatch, match)

		str = "Stamps Jan _2 15:04:05.999 and Jan _2 15:04:05.999999 and Jan _2 15:04:05.999999999"
		match = MatchTimestamp(time.Stamp)(str)

		expectedMatch = Match{
			Template: "Stamps Jan _2 15:04:05.999 and Jan _2 15:04:05.999999 and Jan _2 15:04:05.999999999",
			Patterns: nil,
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("StampMilli", func(t *testing.T) {
		str := "Current timestamp is Jan _2 15:04:05.999"
		match := MatchTimestamp(time.StampMilli)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Jan _2 15:04:05.999"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("StampMicro", func(t *testing.T) {
		str := "Current timestamp is Jan _2 15:04:05.999999"
		match := MatchTimestamp(time.StampMicro)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Jan _2 15:04:05.999999"},
		}

		assert.Equal(t, expectedMatch, match)
	})

	t.Run("StampNano", func(t *testing.T) {
		str := "Current timestamp is Jan _2 15:04:05.000000000"
		match := MatchTimestamp(time.StampNano)(str)

		expectedMatch := Match{
			Template: "Current timestamp is %s",
			Patterns: []string{"Jan _2 15:04:05.000000000"},
		}

		assert.Equal(t, expectedMatch, match)
	})
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

func Test_MatchDaysOfWeek(t *testing.T) {
	str := "Today is Tuesday or tuesday not tUesday"
	actualMatch := MatchDaysOfWeek()(str)
	expectedMatch := Match{Template: "Today is %s or %s not tUesday", Patterns: []string{"Tuesday", "tuesday"}}
	assert.Equal(t, actualMatch, expectedMatch)
	str = "Today is Tuesday or tuesday not tUesday but Tuesday"
	actualMatch = MatchDaysOfWeek()(str)
	expectedMatch = Match{Template: "Today is %s or %s not tUesday but %s", Patterns: []string{"Tuesday", "tuesday", "Tuesday"}}
	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_MatchEmail(t *testing.T) {
	str := "I am <foo@bar.com> and testing to send to dev@test"
	actualMatch := MatchEmail()(str)
	expectedMatch := Match{
		Template: "I am <%s> and testing to send to dev@test",
		Patterns: []string{"foo@bar.com"},
	}
	assert.Equal(t, expectedMatch, actualMatch)

	str = "I am <foo@bar.com> and testing to send to john@doe.io"
	actualMatch = MatchEmail()(str)
	expectedMatch = Match{
		Template: "I am <%s> and testing to send to %s",
		Patterns: []string{"foo@bar.com", "john@doe.io"},
	}
	assert.Equal(t, expectedMatch, actualMatch)
}

func Test_findPatternMatchIndexes(t *testing.T) {

	t.Parallel()

	t.Run("Single Pattern", func(t *testing.T) {
		str := "I scream, you all scream, we all scream for ice cream."
		patterns := []string{"scream"}

		actual := findPatternMatchIndexes(str, patterns)
		expected := map[int]string{
			2:"scream",
			18: "scream",
			33: "scream",
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("Multiple Patterns", func(t *testing.T) {
		str := "I scream, you all scream, we all scream for ice cream."
		patterns := []string{"scream", "ice", "cream"}

		actual := findPatternMatchIndexes(str, patterns)
		expected := map[int]string{
			2:"scream",
			18: "scream",
			33: "scream",
			44: "ice",
			48: "cream",
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("No pattern occurences", func(t *testing.T) {
		str := "I scream, you all scream, we all scream for ice cream."
		patterns := []string{"pickle"}

		actual := findPatternMatchIndexes(str, patterns)
		expected := map[int]string{}
		assert.Equal(t, expected, actual)
	})

	t.Run("Single Regexp Pattern", func(t *testing.T) {
		str := "I scream, you all scream, we all scream for ice cream."
		patterns := []string{"s[a-zA-Z]+"}

		actual := findPatternMatchIndexes(str, patterns)
		expected := map[int]string{
			2:"scream",
			18: "scream",
			33: "scream",
		}
		assert.Equal(t, expected, actual)
	})
	
	t.Run("Multiple Regexp Patterns", func(t *testing.T) {
		str := "I scream, you all scream, we all scream for ice cream."
		patterns := []string{"s[a-zA-Z]+", "(?:,|\\.)"}

		actual := findPatternMatchIndexes(str, patterns)
		expected := map[int]string{
			2:"scream",
			18: "scream",
			33: "scream",
			8: ",",
			24: ",",
			53: ".",
		}
		assert.Equal(t, expected, actual)
	})

}

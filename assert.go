package marker

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func isFuncsEqual(a, b interface{}) bool {
	av := reflect.ValueOf(&a).Elem()
	bv := reflect.ValueOf(&b).Elem()
	return av.InterfaceData() == bv.InterfaceData()
}

func assertFuncEqual(t *testing.T, expected, actual interface{}) {
	assert.True(t, isFuncsEqual(expected, actual), "MatcherFuncs are not equal")
}

func assertMarkRuleEqual(t *testing.T, expected, actual MarkRule) {
	assert.Equal(t, expected.Color, actual.Color, "Colors are not equal")
	assertFuncEqual(t, expected.Matcher, actual.Matcher)
}

func assertMarkRuleSliceEqual(t *testing.T, expected, actual []MarkRule) {
	require.Len(t, actual, len(expected))
	for i := range expected {
		assertMarkRuleEqual(t, expected[i], actual[i])
	}
}

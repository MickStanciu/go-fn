package strings_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/strings"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	input := []string{"A", "B", "C", "a", "b", "c", "d"}
	assert.True(t, strings.SliceContains(input, "B"))
	assert.True(t, strings.SliceContains(input, "A"))
	assert.True(t, strings.SliceContains(input, "d"))
	assert.False(t, strings.SliceContains(input, "e"))
	assert.False(t, strings.SliceContains(input, "z"))
}

func TestGetOrElse(t *testing.T) {
	assert.EqualValues(t, "Z", strings.GetOrElse("Z", "should not happen"))
	assert.EqualValues(t, "Z", strings.GetOrElse("", "Z"))
}

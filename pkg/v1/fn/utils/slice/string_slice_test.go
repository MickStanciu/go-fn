package slice_test

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	input := []string{"A", "B", "C", "a", "b", "c", "d"}
	s := slice.StringSlice{}.Build(input)

	assert.True(t, s.Contains("B"))
	assert.True(t, s.Contains("A"))
	assert.True(t, s.Contains("d"))
	assert.False(t, s.Contains("e"))
	assert.False(t, s.Contains("z"))
}

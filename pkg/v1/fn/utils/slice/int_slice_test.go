package slice_test

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/slice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSlice_Contains(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := slice.IntSlice{}.Build(input)

	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(5))
}

func TestIntSlice_Map(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := slice.IntSlice{}.Build(input)
	expected := []int{2, 4, 6, 8}

	assert.EqualValues(t, expected, s.Map(func(i int) int {
		return i * 2
	}))
}

func TestIntSlice_Filter(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := slice.IntSlice{}.Build(input)
	expected := []int{2, 4}

	assert.EqualValues(t, expected, s.Filter(func(i int) bool {
		return i%2 == 0
	}))
}

func TestIntSlice_Sum(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := slice.IntSlice{}.Build(input)

	assert.EqualValues(t, 10, s.Sum())
}

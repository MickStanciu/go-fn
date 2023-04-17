package ints_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils/ints"
	"github.com/stretchr/testify/assert"
)

func TestIntSlice_Contains(t *testing.T) {
	input := ints.IntSlice{1, 2, 3, 4}
	assert.True(t, input.Contains(2))
	assert.False(t, input.Contains(5))
}

func TestIntSlice_Map(t *testing.T) {
	input := ints.IntSlice{1, 2, 3, 4}
	expected := ints.IntSlice{2, 4, 6, 8}
	assert.EqualValues(t, expected, input.Map(func(i int) int {
		return i * 2
	}))
}

func TestIntSlice_Filter(t *testing.T) {
	input := ints.IntSlice{1, 2, 3, 4}
	expected := ints.IntSlice{2, 4}
	assert.EqualValues(t, expected, input.Filter(func(i int) bool {
		return i%2 == 0
	}))
}

func TestIntSlice_Sum(t *testing.T) {
	input := ints.IntSlice{1, 2, 3, 4}
	assert.EqualValues(t, 10, input.Sum())
}

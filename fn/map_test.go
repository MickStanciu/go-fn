package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.Map(a, func(i int) int {
		return i * 2
	})
	assert.EqualValues(t, []int{2, 4, 6}, b)
}

func TestTransformSliceBy(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: []int{},
		},
		"when no empty": {
			input:          []int{1, 2, 3},
			expectedOutput: []int{2, 4, 6},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.TransformSliceBy(test.input, func(i int) int {
				return i * 2
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

func TestTransformMapBy(t *testing.T) {
	tests := map[string]struct {
		input          map[string]int
		expectedOutput map[string]int
	}{
		"when empty": {
			input:          map[string]int{},
			expectedOutput: map[string]int{},
		},
		"when no empty": {
			input: map[string]int{
				"a": 1, "b": 2, "c": 3,
			},
			expectedOutput: map[string]int{
				"a": 2, "b": 4, "c": 6,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.TransformMapBy(test.input, func(i int) int {
				return i * 2
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: []int{},
		},
		"when no match": {
			input:          []int{1, 2, 3},
			expectedOutput: []int{},
		},
		"when match": {
			input:          []int{10, 20, 30},
			expectedOutput: []int{20, 30},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Filter(test.input, func(i int) bool {
				return i > 10
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

func TestFilterSliceBy(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: []int{},
		},
		"when no match": {
			input:          []int{1, 2, 3},
			expectedOutput: []int{},
		},
		"when match": {
			input:          []int{10, 20, 30},
			expectedOutput: []int{20, 30},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.FilterSliceBy(test.input, func(i int) bool {
				return i > 10
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

func TestFilterMapBy(t *testing.T) {
	tests := map[string]struct {
		input          map[string]int
		expectedOutput map[string]int
	}{
		"when empty": {
			input:          map[string]int{},
			expectedOutput: map[string]int{},
		},
		"when no match": {
			input: map[string]int{
				"a": 1, "b": 2, "c": 3,
			},
			expectedOutput: map[string]int{},
		},
		"when match": {
			input: map[string]int{
				"a": 10, "b": 20, "c": 30,
			},
			expectedOutput: map[string]int{
				"b": 20, "c": 30,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.FilterMapBy(test.input, func(i int) bool {
				return i > 10
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

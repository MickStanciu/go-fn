package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

	"github.com/stretchr/testify/assert"
)

func TestGetOrElse(t *testing.T) {
	tests := map[string]struct {
		a              int
		b              int
		expectedOutput int
	}{
		"when matches": {
			a: 10, b: 20,
			expectedOutput: 10,
		},
		"when not matches": {
			a: 16, b: 20,
			expectedOutput: 20,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.GetOrElse(test.a, test.b, func(i int) bool {
				return i < 15
			})
			assert.Equal(t, test.expectedOutput, result)

		})
	}
}

func TestAny(t *testing.T) {
	sample := []string{"A", "B", "C"}
	assert.True(t, fn.Any[string](sample, func(s string) bool {
		return s == "A"
	}))
	assert.False(t, fn.Any[string](sample, func(s string) bool {
		return s == "Z"
	}))
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	evens := fn.Filter(input, func(i int) bool {
		return i%2 == 0
	})
	assert.EqualValues(t, []int{2, 4, 6, 8}, evens)
}

func TestFilterRight(t *testing.T) {
	tests := map[string]struct {
		input          []string
		expectedOutput []string
	}{
		"when empty": {
			input:          []string{},
			expectedOutput: []string{},
		},
		"when no match": {
			input:          []string{"A", "B", "C"},
			expectedOutput: []string{"A", "B", "C"},
		},
		"when match 1": {
			input:          []string{"A", "B", "C", "D", "X", "E", "F", "X"},
			expectedOutput: []string{"A", "B", "C", "D", "X", "E", "F"},
		},
		"when match multiple": {
			input:          []string{"A", "B", "C", "D", "X", "X", "X", "X"},
			expectedOutput: []string{"A", "B", "C", "D"},
		},
		"when match all": {
			input:          []string{"X", "X", "X", "X"},
			expectedOutput: []string{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.FilterRight(test.input, func(s string) bool {
				return s == "X"
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

func TestAll(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput bool
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: false,
		},
		"when all": {
			input:          []int{3, 3, 3},
			expectedOutput: true,
		},
		"when not all": {
			input:          []int{3, 1, 3},
			expectedOutput: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.All(test.input, func(i int) bool {
				return i == 3
			})
			assert.Equal(t, test.expectedOutput, result)

		})
	}
}

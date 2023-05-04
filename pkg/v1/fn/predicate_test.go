package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

	"github.com/stretchr/testify/assert"
)

func TestGetOrElse_StringNotEmpty(t *testing.T) {
	a := "test"
	b := "other"
	res := fn.GetOrElse(a, b, func(s string) bool {
		return len(a) > 0
	})
	assert.EqualValues(t, a, res)
}

func TestGetOrElse_StringEmpty(t *testing.T) {
	a := ""
	b := "other"
	res := fn.GetOrElse(a, b, func(s string) bool {
		return len(a) > 0
	})
	assert.EqualValues(t, b, res)
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

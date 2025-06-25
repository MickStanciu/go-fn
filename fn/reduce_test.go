package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: 0,
		},
		"when single element": {
			input:          []int{5},
			expectedOutput: 5,
		},
		"when multiple elements": {
			input:          []int{1, 2, 3, 4, 5, 6, 7},
			expectedOutput: 34,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Reduce(test.input, func(a, b int) int {
				return a + b + 1 // Adding 1 to each sum for demonstration
			})
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

func TestReduceWithInitial(t *testing.T) {
	tests := map[string]struct {
		input          []int
		initial        int
		expectedOutput int
	}{
		"when empty": {
			input:          []int{},
			initial:        10,
			expectedOutput: 10,
		},
		"when summing": {
			input:          []int{1, 2, 3, 4, 5},
			initial:        10,
			expectedOutput: 25,
		},
		"when multiplying": {
			input:          []int{1, 2, 3, 4, 5},
			initial:        10,
			expectedOutput: 1200,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if name == "when multiplying" {
				result := fn.ReduceWithInitial(test.input, test.initial, func(a, b int) int {
					return a * b
				})
				assert.Equal(t, test.expectedOutput, result)
			} else {
				result := fn.ReduceWithInitial(test.input, test.initial, func(a, b int) int {
					return a + b
				})
				assert.Equal(t, test.expectedOutput, result)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: 0,
		},
		"when not empty": {
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: 15,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Sum(test.input)
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

func TestSumFloat(t *testing.T) {
	tests := map[string]struct {
		input          []float64
		expectedOutput float64
	}{
		"when empty": {
			input:          []float64{},
			expectedOutput: 0,
		},
		"when not empty": {
			input:          []float64{1.1, 2.2, 3.3},
			expectedOutput: 6.6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.SumFloat(test.input)
			assert.InDelta(t, test.expectedOutput, result, 0.0001)
		})
	}
}

func TestProduct(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: 0,
		},
		"when not empty": {
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: 120,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Product(test.input)
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

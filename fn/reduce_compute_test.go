package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestReduceAndCompute(t *testing.T) {
	type cat struct {
		weight int
	}

	tests := map[string]struct {
		input          []cat
		expectedOutput int
	}{
		"when empty": {
			input:          []cat{},
			expectedOutput: 0,
		},
		"when not empty": {
			input: []cat{
				{weight: 1},
				{weight: 2},
				{weight: 3},
				{weight: 4},
			},
			expectedOutput: 10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.ReduceAndCompute(test.input, func(a cat, b int) int {
				return a.weight + b
			})
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

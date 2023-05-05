package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

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
		"when not empty": {
			input:          []int{1, 2, 3, 4, 5, 6, 7},
			expectedOutput: 34,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Reduce(test.input, func(a, b int) int {
				return a + b + 1
			})
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

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
			expectedOutput: nil,
		},
		"when no match": {
			input:          []int{1, 2, 3},
			expectedOutput: nil,
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

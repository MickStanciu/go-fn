package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

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

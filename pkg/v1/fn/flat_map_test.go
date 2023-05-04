package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedOutput []int
	}{
		"when empty": {
			input:          []int{},
			expectedOutput: nil,
		},
		"when not empty": {
			input:          []int{1, 2, 3},
			expectedOutput: []int{0, 0, 1, 0, 1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.FlatMap(test.input, func(i int) []int {
				var out []int
				for j := 0; j < i; j++ {
					out = append(out, j)
				}
				return out
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

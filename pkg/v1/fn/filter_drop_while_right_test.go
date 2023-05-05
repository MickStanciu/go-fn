package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
)

func TestDropWhileRight(t *testing.T) {
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
			result := fn.DropWhileRight(test.input, func(s string) bool {
				return s == "X"
			})
			assert.EqualValues(t, test.expectedOutput, result)
		})
	}
}

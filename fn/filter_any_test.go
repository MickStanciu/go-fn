package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	tests := map[string]struct {
		input          []string
		expectedOutput bool
	}{
		"when empty": {
			input:          []string{},
			expectedOutput: false,
		},
		"when no match": {
			input:          []string{"A", "B", "C"},
			expectedOutput: false,
		},
		"when match": {
			input:          []string{"A", "B", "C", "D"},
			expectedOutput: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Any(test.input, func(s string) bool {
				return s == "D"
			})
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

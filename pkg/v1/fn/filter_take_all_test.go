package fn_test

import (
	"strings"
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
)

func TestTakeAll(t *testing.T) {
	tests := map[string]struct {
		input          []string
		expectedOutput []string
	}{
		"when empty": {
			input:          []string{},
			expectedOutput: nil,
		},
		"when no match": {
			input:          []string{"A", "B", "C"},
			expectedOutput: nil,
		},
		"when match": {
			input:          []string{"A", "B", "C", "D", "DEF"},
			expectedOutput: []string{"D", "DEF"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.TakeAll(test.input, func(s string) bool {
				return strings.HasPrefix(s, "D")
			})
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

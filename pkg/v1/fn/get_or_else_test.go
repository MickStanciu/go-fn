package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
)

func TestGetOrElse(t *testing.T) {
	tests := map[string]struct {
		a              int
		b              int
		expectedOutput int
	}{
		"when matches": {
			a: 10, b: 20,
			expectedOutput: 10,
		},
		"when not matches": {
			a: 16, b: 20,
			expectedOutput: 20,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.GetOrElse(test.a, test.b, func(i int) bool {
				return i < 15
			})
			assert.Equal(t, test.expectedOutput, result)

		})
	}
}

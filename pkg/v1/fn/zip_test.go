package fn_test

import (
	"fmt"
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	tests := map[string]struct {
		a              []string
		b              []int
		expectedOutput []string
	}{
		"will zip two slices": {
			a:              []string{"a", "b", "c"},
			b:              []int{1, 2, 3},
			expectedOutput: []string{"a-1", "b-2", "c-3"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Zip(test.a, test.b, func(a string, b int) string {
				return fmt.Sprintf("%s-%d", a, b)
			})
			assert.Equal(t, test.expectedOutput, result)

		})
	}
}

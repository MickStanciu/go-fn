package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/fn"
	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	tests := map[string]struct {
		input          []int
		chunkSize      int
		expectedOutput [][]int
	}{
		"when empty": {
			input:          []int{},
			chunkSize:      2,
			expectedOutput: [][]int{},
		},
		"when invalid chunk size": {
			input:          []int{1, 2, 3},
			chunkSize:      0,
			expectedOutput: [][]int{},
		},
		"when evenly divisible": {
			input:          []int{1, 2, 3, 4},
			chunkSize:      2,
			expectedOutput: [][]int{{1, 2}, {3, 4}},
		},
		"when not evenly divisible": {
			input:          []int{1, 2, 3, 4, 5},
			chunkSize:      2,
			expectedOutput: [][]int{{1, 2}, {3, 4}, {5}},
		},
		"when chunk size equals array length": {
			input:          []int{1, 2, 3},
			chunkSize:      3,
			expectedOutput: [][]int{{1, 2, 3}},
		},
		"when chunk size larger than array length": {
			input:          []int{1, 2, 3},
			chunkSize:      5,
			expectedOutput: [][]int{{1, 2, 3}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fn.Chunk(test.input, test.chunkSize)
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

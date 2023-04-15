package fn_test

import (
	"github.com/MickStanciu/go-fn/v1/fn"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce_Integers(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := fn.Reduce(a, func(a, b int) int {
		return a + b + 1
	})
	assert.EqualValues(t, 34, b)
}

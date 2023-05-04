package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.Map(a, func(i int) int {
		return i * 2
	})
	assert.EqualValues(t, []int{2, 4, 6}, b)
}

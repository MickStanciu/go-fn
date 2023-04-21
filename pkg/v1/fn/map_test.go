package fn_test

import (
	"fmt"
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

	"github.com/stretchr/testify/assert"
)

func TestFmap_Strings(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := fn.Fmap(a, func(a string) string {
		return a + "*"
	})
	assert.EqualValues(t, []string{"a*", "b*", "c*"}, b)
}

func TestFmap_Ints(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.Fmap(a, func(a int) int {
		return a * 2
	})
	assert.EqualValues(t, []int{2, 4, 6}, b)
}

func TestFmap_Any(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.Fmap(a, func(a int) string {
		return fmt.Sprintf("$%d", a)
	})
	assert.EqualValues(t, []string{"$1", "$2", "$3"}, b)
}

func TestMap_Ints(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.Map(a, func(i int) int {
		return i * 2
	})
	assert.EqualValues(t, []int{2, 4, 6}, b)
}

func TestFlatMap_Ints(t *testing.T) {
	a := []int{1, 2, 3}
	b := fn.FlatMap(a, func(i int) []int {
		var out []int
		for j := 0; j < i; j++ {
			out = append(out, j)
		}
		return out
	})
	assert.EqualValues(t, []int{0, 0, 1, 0, 1, 2}, b)
}

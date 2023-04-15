package fn_test

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSwap_BothIntegers(t *testing.T) {
	x1, x2 := fn.Swap(4, 5)
	assert.EqualValues(t, 5, x1)
	assert.EqualValues(t, 4, x2)
}

func TestSwap_DifferentTypes(t *testing.T) {
	type human struct {
		age int
	}

	george := human{
		age: 41,
	}

	h, s := fn.Swap("Yo", george)
	assert.EqualValues(t, reflect.TypeOf(george), reflect.TypeOf(h))
	assert.EqualValues(t, 41, h.age)
	assert.EqualValues(t, s, "Yo")
}

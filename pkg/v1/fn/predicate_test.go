package fn_test

import (
	"testing"

	"github.com/MickStanciu/go-fn/pkg/v1/fn"

	"github.com/stretchr/testify/assert"
)

func TestGetOrElse_StringNotEmpty(t *testing.T) {
	a := "test"
	b := "other"
	res := fn.GetOrElse(a, b, func(s string) bool {
		return len(a) > 0
	})
	assert.EqualValues(t, a, res)
}

func TestGetOrElse_StringEmpty(t *testing.T) {
	a := ""
	b := "other"
	res := fn.GetOrElse(a, b, func(s string) bool {
		return len(a) > 0
	})
	assert.EqualValues(t, b, res)
}

func TestAny(t *testing.T) {
	sample := []string{"A", "B", "C"}
	assert.True(t, fn.Any[string](sample, func(s string) bool {
		return s == "A"
	}))
	assert.False(t, fn.Any[string](sample, func(s string) bool {
		return s == "Z"
	}))
}

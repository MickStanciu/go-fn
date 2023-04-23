package slice

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
)

type Slice[T any] []T

// GetSlice - returns the slice
func (s Slice[T]) GetSlice() Slice[T] {
	return s
}

// Head - returns top of the slice or zero-value for the given type
func (s Slice[T]) Head() T {
	if len(s) > 0 {
		return s[0]
	}
	return *new(T)
}

// Map - returns a new slice where every element has been transformed by the map function
func (s Slice[T]) Map(f fn.MapFn[T, T]) Slice[T] {
	return fn.Map(s, f)
}

// Filter - returns a new slice filtered by the predicate function
func (s Slice[T]) Filter(f fn.Predicate[T]) Slice[T] {
	return fn.Filter(s, f)
}

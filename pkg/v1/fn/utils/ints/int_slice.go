package ints

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
)

type IntSlice []int

// Map - returns a new slice where every element has been transformed by the map function
func (i IntSlice) Map(f fn.MapFn[int, int]) IntSlice {
	return fn.Map(i, f)
}

// Filter - returns a new slice filtered by the predicate function
func (i IntSlice) Filter(f fn.Predicate[int]) IntSlice {
	return fn.Filter(i, f)
}

// Sum - sums the slice
func (i IntSlice) Sum() int {
	return fn.Sum(i)
}

// Contains - returns `true` if the slice contains the element `n`
func (i IntSlice) Contains(n int) bool {
	return fn.Any(i, func(elem int) bool {
		return elem == n
	})
}

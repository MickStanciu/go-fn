package ints

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
)

type IntSlice []int

func (i IntSlice) Map(f func(i int) int) IntSlice {
	return fn.Map(i, f)
}

func (i IntSlice) Filter(f func(i int) bool) IntSlice {
	return fn.Filter(i, f)
}

func (i IntSlice) Sum() int {
	return fn.Sum(i)
}

func (i IntSlice) Contains(n int) bool {
	return fn.Any(i, func(elem int) bool {
		return elem == n
	})
}

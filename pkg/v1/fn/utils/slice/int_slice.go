package slice

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
)

type IntSlice struct {
	Slice[int]
}

// Build - construct a StringSlice with values
func (s IntSlice) Build(input []int) IntSlice {
	return IntSlice{input}
}

// Sum - sums the slice
func (i IntSlice) Sum() int {
	return fn.Sum(i.GetSlice())
}

// Contains - returns `true` if the slice contains the element `n`
func (i IntSlice) Contains(n int) bool {
	return fn.Any(i.GetSlice(), func(elem int) bool {
		return elem == n
	})
}

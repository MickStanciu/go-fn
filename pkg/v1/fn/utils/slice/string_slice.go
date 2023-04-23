package slice

import (
	"github.com/MickStanciu/go-fn/pkg/v1/fn"
)

type StringSlice struct {
	Slice[string]
}

// Build - construct a StringSlice with values
func (s StringSlice) Build(input []string) StringSlice {
	return StringSlice{input}
}

// Contains - returns true if `n` is in the collection
func (s StringSlice) Contains(n string) bool {
	return fn.Any(s.GetSlice(), func(elem string) bool {
		return elem == n
	})
}

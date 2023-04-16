package strings

import "github.com/MickStanciu/go-fn/pkg/v1/fn"

type StringSlice []string

func (s StringSlice) Contains(n string) bool {
	return fn.Any(s, func(elem string) bool {
		return elem == n
	})
}

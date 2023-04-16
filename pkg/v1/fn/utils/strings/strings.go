package strings

import "github.com/MickStanciu/go-fn/pkg/v1/fn"

// GetOrElse - will return 'a' if the string is not empty or 'b'
func GetOrElse(a, b string) string {
	return fn.GetOrElse(a, b, func(s string) bool {
		return len(s) > 0
	})
}

// SliceContains - returns true if `n` is in the collection `col`
func SliceContains(col []string, n string) bool {
	return StringSlice(col).Contains(n)
}

package fn

// Filter - filters a slice of types T, using a predicate function
// Returns an empty slice if no match, or the filtered collection
// Deprecated: use the FilterSliceBy instead. No contract change.
func Filter[T any](input []T, p Predicate[T]) []T {
	return FilterSliceBy(input, p)
}

func FilterSliceBy[T any](input []T, p Predicate[T]) []T {
	var out = make([]T, 0)

	for _, element := range input {
		if p(element) {
			out = append(out, element)
		}
	}

	return out
}

// FilterMapBy - filters a map[KEY]T, using a predicate function
// Returns an empty map if no match, or the filtered collection
func FilterMapBy[KEY string, U any](input map[KEY]U, p Predicate[U]) map[KEY]U {
	var out = make(map[KEY]U)

	for key, element := range input {
		if p(element) {
			out[key] = element
		}
	}

	return out
}

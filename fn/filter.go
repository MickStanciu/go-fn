package fn

// Filter - filters a slice of types T, using a predicate function
// Returns an empty slice if no match, or the filtered collection
// Deprecated: use the FilterSliceBy instead. No contract change.
func Filter[T any](input []T, p Predicate[T]) []T {
	return FilterSliceBy(input, p)
}

func FilterSliceBy[T any](input []T, p Predicate[T]) []T {
	if len(input) == 0 {
		return []T{}
	}

	// Pre-allocate with a reasonable initial capacity
	// In worst case (all elements match), we need len(input)
	out := make([]T, 0, len(input))

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
	if len(input) == 0 {
		return map[KEY]U{}
	}

	// Pre-allocate with a reasonable initial capacity
	// In worst case (all elements match), we need len(input)
	out := make(map[KEY]U, len(input))

	for key, element := range input {
		if p(element) {
			out[key] = element
		}
	}

	return out
}

package fn

// TakeAll - filters a collection of type T, using a predicate function,
// returning the elements which satisfy the predicate
func TakeAll[T any](input []T, p Predicate[T]) []T {
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

	// Return nil if no elements matched the predicate
	if len(out) == 0 {
		return []T{}
	}

	return out
}

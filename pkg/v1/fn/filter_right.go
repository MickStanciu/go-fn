package fn

// FilterRight - filters a collection of type T, using a predicate function,
// by removing the right-most elements which satisfy the predicate, while preserving the order
func FilterRight[T any](input []T, p Predicate[T]) []T {
	ln := len(input)
	found := 0
	for i := ln - 1; i >= 0; i-- {
		if !p(input[i]) {
			break
		}
		found++
	}

	return input[:ln-found]
}

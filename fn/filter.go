package fn

// Filter - filters a collection of type T, using a predicate function
func Filter[T any](input []T, p Predicate[T]) []T {
	var out []T

	for _, element := range input {
		if p(element) {
			out = append(out, element)
		}
	}

	return out
}

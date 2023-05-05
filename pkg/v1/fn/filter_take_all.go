package fn

// TakeAll - filters a collection of type T, using a predicate function,
// returning the elements which satisfy the predicate
func TakeAll[T any](input []T, p Predicate[T]) []T {
	var out []T

	for _, element := range input {
		if p(element) {
			out = append(out, element)
		}
	}

	return out
}

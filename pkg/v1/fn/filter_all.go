package fn

// All - returns true if all elements are satisfying the predicate function
func All[T any](input []T, p Predicate[T]) bool {
	if len(input) == 0 {
		return false
	}

	for _, element := range input {
		if !p(element) {
			return false
		}
	}
	return true
}

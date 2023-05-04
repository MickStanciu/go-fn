package fn

// Any - returns true if one element satisfies the predicate function
func Any[T any](input []T, p Predicate[T]) bool {
	for _, element := range input {
		if p(element) {
			return true
		}
	}
	return false
}

package fn

// GetOrElse - will return input if predicate is satisfied or other
func GetOrElse[T any](input T, other T, p Predicate[T]) T {
	if p(input) {
		return input
	}
	return other
}

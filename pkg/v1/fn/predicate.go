package fn

// Predicate - statement which can be evaluated to true/false
type Predicate[T any] func(T) bool

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

// Any - returns true if one element satisfies the predicate function
func Any[T any](input []T, p Predicate[T]) bool {
	for _, element := range input {
		if p(element) {
			return true
		}
	}
	return false
}

// All - returns true if all elements are satisfying the predicate function
func All[T any](input []T, p Predicate[T]) bool {
	for _, element := range input {
		if !p(element) {
			return false
		}
	}
	return true
}

// GetOrElse - will return input if predicate is satisfied or other
func GetOrElse[T any](input T, other T, p Predicate[T]) T {
	if p(input) {
		return input
	}
	return other
}

// FilterRight - filters a collection of type T, using a predicate function,
// by removing the right-most elements which satisfy the predicate
func FilterRight[T any](input []T, p Predicate[T]) []T {
	ln := len(input)
	if ln == 0 {
		return input
	}

	found := 0
	for i := ln - 1; i >= 0; i-- {
		if !p(input[i]) {
			break
		}
		found++
	}

	return input[:ln-found]
}

package fn

type ReduceFn[T any] func(a, b T) T

// Reduce - will fold a collection
func Reduce[T any](input []T, fn ReduceFn[T]) T {
	if len(input) == 0 {
		return *new(T)
	}

	result := input[0]
	for _, element := range input[1:] {
		result = fn(result, element)
	}
	return result
}

// Sum - will add integers
func Sum[T ~int](in []T) T {
	return Reduce(in, func(a, b T) T {
		return a + b
	})
}

package fn

type MapFn[T any] func(T) T
type FlatMapFn[T any] func(T) []T
type FMapFn[A, B any] func(A) B

// Fmap - applies a transformation function A -> B to each element of type A
func Fmap[A, B any](a []A, fn FMapFn[A, B]) []B {
	b := make([]B, len(a))
	for i, elem := range a {
		b[i] = fn(elem)
	}
	return b
}

// Map - applies a transformation function T -> T to each element of type T
func Map[T any](input []T, fn MapFn[T]) []T {
	output := make([]T, len(input))
	for i, element := range input {
		output[i] = fn(element)
	}
	return output
}

// FlatMap - applies a transformation function from T to []T
func FlatMap[T any](input []T, fn FlatMapFn[T]) []T {
	var output []T
	for _, element := range input {
		elems := fn(element)
		output = append(output, elems...)
	}
	return output
}
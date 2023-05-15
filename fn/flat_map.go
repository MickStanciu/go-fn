package fn

// FlatMap - applies a transformation function from T to []T to each element of type T
func FlatMap[T any](input []T, fn FlatMapFn[T]) []T {
	var output []T
	for _, element := range input {
		elems := fn(element)
		output = append(output, elems...)
	}
	return output
}

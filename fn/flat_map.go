package fn

// FlatMap - applies a transformation function from T to []T to each element of type T
func FlatMap[T any](input []T, fn FlatMapFn[T]) []T {
	if len(input) == 0 {
		return []T{}
	}

	// Estimate initial capacity to reduce allocations
	// First pass: calculate total size
	totalSize := 0
	for _, element := range input {
		totalSize += len(fn(element))
	}

	// Second pass: populate the pre-allocated slice
	output := make([]T, 0, totalSize)
	for _, element := range input {
		elems := fn(element)
		output = append(output, elems...)
	}
	return output
}

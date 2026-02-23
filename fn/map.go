package fn

// Map - applies a transformation function A -> B to each element of type A
// Deprecated: use the TransformSliceBy instead. No contract change.
func Map[A, B any](input []A, fn MapFn[A, B]) []B {
	return TransformSliceBy(input, fn)
}

// TransformSliceBy - applies a transformation function A -> B to each element of type A
// This function handles empty input gracefully and pre-allocates the result slice
func TransformSliceBy[A, B any](input []A, fn MapFn[A, B]) []B {
	if len(input) == 0 {
		return []B{}
	}

	output := make([]B, len(input))
	for i, element := range input {
		output[i] = fn(element)
	}

	return output
}

// TransformMapBy - applies a transformation function A -> B to each element of type A
// This function handles empty input gracefully and pre-allocates the result map
func TransformMapBy[A, B any](input map[string]A, fn MapFn[A, B]) map[string]B {
	if len(input) == 0 {
		return map[string]B{}
	}

	// Pre-allocate with the exact size needed
	out := make(map[string]B, len(input))

	for key, element := range input {
		out[key] = fn(element)
	}

	return out
}

// ConvertMapToSliceBy - converts a map[K]R to a slice of R values
func ConvertMapToSliceBy[K string, R any](input map[K]R) []R {
	result := make([]R, len(input))
	idx := 0
	for _, val := range input {
		result[idx] = val
		idx++
	}

	return result
}

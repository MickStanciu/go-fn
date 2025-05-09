package fn

// Map - applies a transformation function A -> B to each element of type A
// Deprecated: use the TransformSliceBy instead. No contract change.
func Map[A, B any](input []A, fn MapFn[A, B]) []B {
	return TransformSliceBy(input, fn)
}

func TransformSliceBy[A, B any](input []A, fn MapFn[A, B]) []B {
	var output = make([]B, len(input))

	for i, element := range input {
		output[i] = fn(element)
	}

	return output
}

// TransformMapBy - applies a transformation function A -> B to each element of type A
func TransformMapBy[A, B any](input map[string]A, fn MapFn[A, B]) map[string]B {
	var out = make(map[string]B)

	for key, element := range input {
		out[key] = fn(element)
	}

	return out
}

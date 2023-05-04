package fn

// Map - applies a transformation function A -> B to each element of type A
func Map[A, B any](input []A, fn MapFn[A, B]) []B {
	output := make([]B, len(input))
	for i, element := range input {
		output[i] = fn(element)
	}
	return output
}

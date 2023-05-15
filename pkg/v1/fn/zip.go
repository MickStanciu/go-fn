package fn

// Zip - applies a transformation function A -> B -> C to each element of type A and B
func Zip[A, B, C any](a []A, b []B, fn func(A, B) C) []C {
	if len(a) != len(b) {
		panic("cannot zip slices of different lengths")
	}

	result := make([]C, len(a))
	for i := range a {
		result[i] = fn(a[i], b[i])
	}
	return result
}

package fn

// Reduce - applies a reducer function to each element in the slice,
// accumulating the results
// Note: For empty slices, returns the zero value of type T
func Reduce[T any](input []T, fn ReduceFn[T]) T {
	if len(input) == 0 {
		return *new(T) // Return zero value for empty slices
	}

	result := input[0]
	for i := 1; i < len(input); i++ {
		result = fn(result, input[i])
	}

	return result
}

// ReduceWithInitial - applies a reducer function to each element in the slice,
// starting with the provided initial value
func ReduceWithInitial[T any](input []T, initial T, fn ReduceFn[T]) T {
	result := initial

	for _, item := range input {
		result = fn(result, item)
	}

	return result
}

// Sum - will add integers
func Sum[T ~int](in []T) T {
	return Reduce(in, func(a, b T) T {
		return a + b
	})
}

// SumFloat - will add floating point numbers
func SumFloat[T ~float64](in []T) T {
	return Reduce(in, func(a, b T) T {
		return a + b
	})
}

// Product - will multiply integers
func Product[T ~int](in []T) T {
	if len(in) == 0 {
		return 0
	}

	return ReduceWithInitial(in[1:], in[0], func(a, b T) T {
		return a * b
	})
}

// ProductFloat - will multiply floating point numbers
func ProductFloat[T ~float64](in []T) T {
	if len(in) == 0 {
		return 0
	}

	return ReduceWithInitial(in[1:], in[0], func(a, b T) T {
		return a * b
	})
}

package fn

type ComputeFn[T, R any] func(T, R) R
type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// ReduceAndCompute - will fold a collection to a Number
func ReduceAndCompute[T any, R Number](input []T, fn ComputeFn[T, R]) R {
	result := R(0)

	if len(input) == 0 {
		return result
	}

	for _, element := range input {
		result = fn(element, result)
	}
	return result
}

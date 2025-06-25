package fn

// Chunk - splits a slice into chunks of specified size
// The last chunk may contain fewer elements if the slice length is not divisible by the chunk size
func Chunk[T any](input []T, size int) [][]T {
	if len(input) == 0 || size <= 0 {
		return [][]T{}
	}

	chunksCount := (len(input) + size - 1) / size // Ceiling division
	result := make([][]T, 0, chunksCount)

	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}

		chunk := input[i:end]
		result = append(result, chunk)
	}

	return result
}

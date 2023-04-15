package fn

// Swap - swaps two values around
func Swap[A, B any](a A, b B) (B, A) {
	return b, a
}

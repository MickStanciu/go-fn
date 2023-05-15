package fn

// ToPtr - returns a pointer to the input
func ToPtr[T any](in T) *T {
	return &in
}

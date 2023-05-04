package fn

// Predicate - statement which can be evaluated to true/false
type Predicate[T any] func(T) bool

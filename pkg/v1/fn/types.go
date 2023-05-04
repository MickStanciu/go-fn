package fn

// Predicate - statement which can be evaluated to true/false
type Predicate[T any] func(T) bool

// MapFn - map function that takes A, and a transformation function A -> B, returns B
type MapFn[A, B any] func(A) B

// FlatMapFn - map function that transforms A -> []B, returns []B
type FlatMapFn[T any] func(T) []T

# go-fn

## How to install
`go get github.com/MickStanciu/go-fn`

## What is it?
This is a collection of useful small Go functions that are using generics

## How to use
### Filter
statement which can be evaluated to true/false
```go
type Predicate[T any] func(T) bool
func Filter[T any](input []T, p Predicate[T]) []T
```

```go
input := []int{1, 2, 3, 4, 5, 6, 7, 8}
evens := fn.Filter(input, func(i int) bool {
    return i%2 == 0
})
```

### Map
Applies a transformation function A -> B to each element of type A
```go
type MapFn[A, B any] func(A) B
func Map[A, B any](input []A, fn MapFn[A, B]) []B
```

```go
a := []string{"George", "Maria", "John"}
fn.Fmap(a, func(a string) string {
    return "hello " + a 
})
```




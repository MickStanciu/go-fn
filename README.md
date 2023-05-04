# go-fn

## How to install
`go get github.com/MickStanciu/go-fn`

## What is it?
This is a collection of useful small Go functions that are using generics

## How to use
### Filters
statement which can be evaluated to true/false
```go
type Predicate[T any] func(T) bool
```

#### Filter
will filter a collection based on the provided predicate
```go
func Filter[T any](input []T, p Predicate[T]) []T
```

Example:
```go
input := []int{1, 2, 3, 4, 5, 6, 7, 8}
evens := fn.Filter(input, func(i int) bool {
    return i%2 == 0
})
```

#### Any
returns true if one element satisfies the predicate function
```go
func Any[T any](input []T, p Predicate[T]) bool
```

Example:
```go
sample := []string{"A", "B", "C"}
result := fn.Any[string](sample, func(s string) bool {
    return s == "A"
})
```

#### All
return true if all elements are satisfying the predicate function
```go
func All[T any](input []T, p Predicate[T]) bool
```

Example:
```go
sample := []string{"A", "A", "A"}
result := fn.All[string](sample, func(s string) bool {
    return s == "A"
})
```

#### GetOrElse
will return the input if the predicate is satisfied, otherwise will return the `else`
```go
func GetOrElse[T any](input T, other T, p Predicate[T]) T
```

Example:
```go
res := fn.GetOrElse(10, 20, func(i int) bool {
    return i > 15
})
```

#### FilterRight
filters a collection of type T, using a predicate function, by removing the right-most elements which satisfy the predicate, while preserving the order

```go
func FilterRight[T any](input []T, p Predicate[T]) []T
```

Example: 
```go
result := fn.FilterRight([]string{"A", "B", "C"}, func(s string) bool {
    return s == "C"
})
```

### Maps
transformation functions
```go
type MapFn[A, B any] func(A) B
````

#### Map
applies a transformation function A -> B to each element of type A
```go
func Map[A, B any](input []A, fn MapFn[A, B]) []B
```

```go
a := []string{"George", "Maria", "John"}
fn.Fmap(a, func(a string) string {
    return "Hello " + a 
})
```




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
evens := fn.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(i int) bool {
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
result := fn.Any[string]([]string{"A", "B", "C"}, func(s string) bool {
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
result := fn.All[string]([]string{"A", "A", "A"}, func(s string) bool {
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

#### DropWhileRight
filters a collection of type T, using a predicate function, by removing the right-most elements which satisfy the predicate, while preserving the order

```go
func DropWhileRight[T any](input []T, p Predicate[T]) []T
```

Example: 
```go
result := fn.DropWhileRight([]string{"A", "B", "C"}, func(s string) bool {
    return s == "C"
})
```

#### TakeAll
filters a collection of type T, using a predicate function, returning the elements which satisfy the predicate

```go
func TakeAll[T any](input []T, p Predicate[T]) []T
```

Example:
```go
result := fn.TakeAll([]string{"A","B","C","D","E"}, func(s string) bool {
    return strings.HasPrefix(s, "D")
})
```

#### DeDuplicateList
makes sure the elements in a collection are unique, based on string KEY

```go
func DeduplicateList[T any](elements []*T, pkFun func(element *T) string) []*T
```

Example:
```go
type human struct {
    ID   string
    Name string
    Age  int
}

result := fn.DeduplicateList(tt.input, func(element *human) string {
    return element.ID
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
fn.Fmap([]string{"George", "Maria", "John"}, func(a string) string {
    return "Hello " + a 
})
```

#### FlatMap
applies a transformation function from T to []T to each element of type T
```go
result := fn.FlatMap([]int{1, 2, 3}, func(i int) []int {
    var out []int
    for j := 0; j < i; j++ {
        out = append(out, j)
    }
    return out
})
```

#### Reduce
will fold a collection
```go
func Reduce[T any](input []T, fn ReduceFn[T]) T
```

Example:
```go
fn.Reduce([]int{1, 2, 3}, func(a, b int) int {
    return a + b + 1
})
```
result is `7`


#### Zip
will combine 2 collections
```go
func Zip[A, B, C any](a []A, b []B, fn func(A, B) C) []C
```

Example:
```go
x := []string{"a", "b", "c"}
y := []int{1, 2, 3}
fn.Zip(x, y, func(a string, b int) string {
    return fmt.Sprintf("%s-%d", a, b)
}
```

#### SplitSliceInBatch
will split a slice into batches and then will call the callback function for each batch.
```go
SplitSliceInBatch[T any](size int, collection []T, fn func(batch []T) error) error
```

Example:
```go
err := batch.SplitSliceInBatch(2, []string{"a", "b", "c", "d", "e", "f"},
    func(strings []string) error {
        if strings[0] == "d" {
            return fmt.Errorf("error in processing function")
        }
        return nil
    })
```

#### ParallelExecByKey
will batch execute a given function, where items are identified by a key
```go
func ParallelExecByKey[R any, Key string](ctx context.Context, batchSize int, keys []Key, fn func(ctx context.Context, key Key) (R, error)) (map[Key]R, error)
```

Example:
```go
res, err := batch.ParallelExecByKey(
		ctx, 2,
		[]string{"item_1", "item_2", "item_3", "item_4", "item_5", "item_6"},
		func(ctx context.Context, itemID string) (string, error) {
			return fmt.Sprintf("%s_", itemID), nil
		},
	)
```
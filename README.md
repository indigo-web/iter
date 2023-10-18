# iter
Yet another iterating toolset.

# Why?
I faced the need of an iterating toolset in indigo. However, I failed to find any good solutions (the most popular one is empty-interfaces-based). That's why this package was created

# What is the difference with another packages?
I have no idea. This one is original, and was constructed according to my own taste

# How to use?
Just `go get github.com/indigo-web/iter`

# Documentation
`iter` aims primarly at being generic, so may be used in various cases with various types. Also iterators are considered to be lazy - that's why a main way of interacting with other iterators for producing result sequence are functional tools: `Map` and `Filter`. All them are simply wrapping other iterators, that may be absolutely anything. 

Whole package walks around the basic definition of an iterator:
```go
type Iterator[T any] interface {
	Next() (el T, cont bool)
	Stopped() bool
	Break()
}
```

To instantiate an iterator from primitive, you need to wrap your slice into `Iterator[T]` interface. To proceed it, simply do:
```go
iter.Slice[T](mySlice)
```
...where T is a type of the slice. For example:
```go
iter.Slice([]byte("Hello, world!")) // generic type inferred automatically
```

### Extract
Function `Extract` takes an iterator and a buffer (may be nil), and appends all the values from an iterator into the passed buffer, returning the filled with values version of it

### Map
Map takes another iterator as a first argument, and a map function by itself as a second. A map function is a function, that takes a single input argument, and a single return value. In case of `iter.Map[string, int]`, map function must take a string as an argument, and int as a return value.
For example, let's convert a slice of strings into a slice of integers:
```go
myStrings := []string{"123", "456"}
sliceIterator := iter.Slice[string](myStrings)
ints := iter.Map[string, int](func(from string) (to int) {
  to, _ = strconv.Atoi(from)
  return to
}, sliceIterator)

fmt.Println(iter.Extract(ints, nil))
// Output: [123 456]
```

### Filter
Filter just skips values, if a filter function says "no". For example:
```go
integers := []int{1, 2, 3, 4, 5}
filtered := iter.Filter[int](func(integer int) bool {
  return integer != 3
}, iter.Slice(integers))

fmt.Println(iter.Extract(filtered, nil))
// Output: 1, 2, 4, 5
```

### Reduce
Reduce produces a single value of a sequence. It iterates over the iterator and passes both previous and current values into the reducing function. For example:
```go
integers := []int{1, 2, 3, 4, 5}
sum := iter.Reduce[int](func(a, b int) int {
  return a + b
}, iter.Slice(integers))

fmt.Println(sum)
// Output: 15
```

There are also some pre-defined filter functions in iter/filter package:
- Unique - is a HoF. It takes a buffer, which will store all the unique values, and returns a `FilterFunc[T]`, that can be used as a filter function. It is approximately O(n^2) (may be wrong), but as project is mainly used by indigo, there it appears to be cheaper than a built-in hashmap (up to 10-20 unique elements, that is enough)

# Post-scriptum
The package doesn't support maps yet, as implementation of it won't be that smooth (it isn't possible in a normal way to make a lazy iterating over a map). There're actually a lot of things that I would do in terms of design and API of the library, so be free to open PRs - I'm gonna review them all🎉

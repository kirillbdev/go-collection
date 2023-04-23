go-collection
=======

[![Go Report Card](https://goreportcard.com/badge/github.com/kirillbdev/go-collection)](https://goreportcard.com/report/github.com/kirillbdev/go-collection)

Generic collections library for common usage

### This is a draft. Do not use in production if you found this package.


### Contract
```go
Items() []T // Get all items from collection
Each(func(item T)) // Walk over items and do something
Count() int // Get total items
IsEmpty() bool // Check if collection has no items
```

### Examples
#### Simple collection

```go
ints := collection.NewSimpleCollection([]int{1, 5, 10, -4, 3})
ints.Each(func(item int) {
    fmt.Println(item)
})
```
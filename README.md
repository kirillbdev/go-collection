go-collection
=======

[![Go Report Card](https://goreportcard.com/badge/github.com/kirillbdev/go-collection)](https://goreportcard.com/report/github.com/kirillbdev/go-collection)

Generic based collections library for common usage.

#### This module is on develop stage. Use it in production at one's own risk.


### How to use
Library contains collection interface with next declaration
```go
Items() []T // Get all items from collection
Each(fn func(item T)) // Walk over items and do something
Count() int // Get total items
IsEmpty() bool // Check if collection has no items
Map(fn func(item T) any) Collection[any] // Map each values of collection
```

Use simple implementation called SimpleCollection
```go
ints := collection.NewSimpleCollection([]int{1, 5, 10, -4, 3})
ints.Each(func(item int) {
    fmt.Println(item)
})
```

### Utils
Utils package contains some additional helpers functions
#### utils.Convert
```go
// Map collection and convert to concrete type
c := collection.NewSimpleCollection([]int{1, 2, 3, 0, 5})
mapped := c.Map(func(item int) any {
    return strconv.Itoa(item) + "_as_string"
})
// Now, mapped has some abstract type of collection after mapping. Concrete it.
concrete := utils.Convert[string](mapped)
```
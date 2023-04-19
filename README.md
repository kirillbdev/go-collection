## This is a draft. Do not use in production if you found this package.

### Generic collections library for common usage

### Contract
```go
Items() []T // Get all items from collection
Each(func(item T)) // Walk over items and do something
```

### Examples
#### Simple collection

```go
ints := collection.NewSimpleCollection([]int{1, 5, 10, -4, 3})
ints.Each(func(item int) {
    fmt.Println(item)
})
```
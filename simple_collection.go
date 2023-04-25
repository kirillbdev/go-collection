package collection

type SimpleCollection[T any] struct {
	items []T
}

func NewSimpleCollection[T any](items []T) *SimpleCollection[T] {
	return &SimpleCollection[T]{
		items: items,
	}
}

func (c *SimpleCollection[T]) Items() []T {
	return c.items
}

func (c *SimpleCollection[T]) Each(fn func(item T)) {
	for _, it := range c.items {
		fn(it)
	}
}

func (c *SimpleCollection[T]) Count() int {
	return len(c.items)
}

func (c *SimpleCollection[T]) IsEmpty() bool {
	return len(c.items) == 0
}

func (c *SimpleCollection[T]) Map(fn func(item T) any) Collection[any] {
	var result []interface{}

	for _, it := range c.items {
		result = append(result, fn(it))
	}

	return NewSimpleCollection(result)
}

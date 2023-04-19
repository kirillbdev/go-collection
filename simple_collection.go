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

func (c *SimpleCollection[T]) Each(fn func(item *T)) {
	for _, it := range c.items {
		fn(&it)
	}
}

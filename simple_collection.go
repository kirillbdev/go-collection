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

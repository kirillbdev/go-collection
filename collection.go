package collection

type Collection[T any] interface {
	Items() []T
	Each(func(item T))
}

package collection

type Collection[T any] interface {
	Items() []T
	Each(func(item T))
	Count() int
	IsEmpty() bool
}

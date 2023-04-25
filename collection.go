package collection

type Collection[T any] interface {
	Items() []T
	Each(fn func(item T))
	Count() int
	IsEmpty() bool
	Map(fn func(item T) any) Collection[any]
}

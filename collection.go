package collection

type Collection[T any] interface {
	Items() []T
}

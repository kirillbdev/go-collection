package utils

import (
	"github.com/kirillbdev/go-collection"
)

// Convert converts one collection with abstract type to collection with concrete type
// This function is only helper to convert collection after mapping. It's not performed casting from one type to other.
//
//	c := collection.NewSimpleCollection([]int{1, 2, 3, 0, 5})
//	abstractCollection := c.Map(func(item int) any {
//		return strconv.Itoa(item) + "_str"
//	})
//	concreteCollection := utils.Convert[string](abstractCollection)
func Convert[T any](c collection.Collection[any]) collection.Collection[T] {
	var result []T

	for _, value := range c.Items() {
		result = append(result, value.(T))
	}

	return collection.NewSimpleCollection(result)
}

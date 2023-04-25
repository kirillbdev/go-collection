package collection

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type countTestPair struct {
	items    []int
	expected int
}

type simpleCollectionIsEmptyTestPair struct {
	items   []int
	isEmpty bool
}

type simpleCollectionMapTestPair[T, K any] struct {
	items    Collection[T]
	expected []K
}

func convert[T any](c Collection[interface{}]) Collection[T] {
	var result []T

	for _, value := range c.Items() {
		result = append(result, value.(T))
	}

	return NewSimpleCollection(result)
}

func TestSimpleCollection_Count(t *testing.T) {
	testPairs := []countTestPair{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{3, 4, 5}, 3},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, pair := range testPairs {
		var c Collection[int] = NewSimpleCollection(pair.items)
		assert.Equal(t, pair.expected, c.Count())
	}
}

func TestSimpleCollection_IsEmpty(t *testing.T) {
	testPairs := []simpleCollectionIsEmptyTestPair{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{3, 4, 5}, false},
		{[]int{}, true},
	}

	for _, pair := range testPairs {
		var c Collection[int] = NewSimpleCollection(pair.items)
		assert.Equal(t, pair.isEmpty, c.IsEmpty())
	}
}

func TestSimpleCollection_Map(t *testing.T) {
	// Case 1
	case1 := simpleCollectionMapTestPair[int, string]{
		NewSimpleCollection([]int{1, 2, 3}),
		[]string{"1", "2", "3"},
	}

	result1 := case1.items.Map(func(item int) any {
		return strconv.Itoa(item)
	})

	assert.Equal(t, case1.expected, convert[string](result1).Items())

	// Case 2
	case2 := simpleCollectionMapTestPair[string, string]{
		NewSimpleCollection([]string{"a", "b", "c"}),
		[]string{"a_a", "b_b", "c_c"},
	}

	result2 := case2.items.Map(func(item string) any {
		return item + "_" + item
	})

	assert.Equal(t, case2.expected, convert[string](result2).Items())
}

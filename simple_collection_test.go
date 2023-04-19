package collection

import "testing"

type countTestPair struct {
	items    []int
	expected int
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
		count := c.Count()

		if count != pair.expected {
			t.Error("Expected ", pair.expected, " got ", count)
		}
	}
}

type simpleCollectionIsEmptyTestPair struct {
	items   []int
	isEmpty bool
}

func TestSimpleCollection_IsEmpty(t *testing.T) {
	testPairs := []simpleCollectionIsEmptyTestPair{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{3, 4, 5}, false},
		{[]int{}, true},
	}

	for _, pair := range testPairs {
		var c Collection[int] = NewSimpleCollection(pair.items)
		empty := c.IsEmpty()

		if empty != pair.isEmpty {
			t.Error("Expected ", pair.isEmpty, " got ", empty)
		}
	}
}

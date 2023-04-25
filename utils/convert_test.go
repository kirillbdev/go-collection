package utils

import (
	"github.com/kirillbdev/go-collection"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	items1 := collection.NewSimpleCollection([]int{1, 2, 3, 0, 5})
	expected1 := collection.NewSimpleCollection([]string{"1", "2", "3", "0", "5"})
	result1 := items1.Map(func(item int) any {
		return strconv.Itoa(item)
	})
	assert.Equal(t, expected1, Convert[string](result1))

	items2 := collection.NewSimpleCollection([]int{1, 2, 3, 0, 5})
	expected2 := collection.NewSimpleCollection([]float32{1.0, 2.0, 3.0, 0, 5.0})
	result2 := items2.Map(func(item int) any {
		return float32(item)
	})
	assert.Equal(t, expected2, Convert[float32](result2))
}

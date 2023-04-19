package main

import (
	"fmt"
	"github.com/kirillbdev/go-collection"
)

func main() {
	ints := collection.NewSimpleCollection([]int{1, 5, 10, -4, 3})
	ints.Each(func(item int) {
		fmt.Println(item)
	})
}

package main

import (
	"fmt"
	"slices"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	ch <- t.Value
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	arr1 := make([]int, 10)
	arr2 := make([]int, 10)

	// store values in a slice and sorting it
	for i := 0; i < 10; i++ {
		arr1[i] = <-ch1
		arr2[i] = <-ch2
	}
	slices.Sort(arr1)
	slices.Sort(arr2)

	// determining if they are same
	for i := 0; i < 10; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}

package main

import (
	"fmt"
)

func main() {
	// when slice does not depend on array
	var s []int = []int{1, 2, 3, 4}
	printSlice(s)

	//slice to give zero length
	s = s[:0]
	printSlice(s)

	//slice to extend its length
	s = s[:4] // This should not exceed the initial largest value or initialize it again
	printSlice(s)

	//slice to drop values
	s = s[2:]
	printSlice(s)

	//nil slices
	var a []int
	fmt.Println(a)
	if a == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Println(cap(s), len(s))
}

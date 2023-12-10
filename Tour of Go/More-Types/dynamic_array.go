package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 6)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Println(s, len(s), cap(s))
}

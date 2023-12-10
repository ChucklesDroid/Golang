package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 2, 4, 5}
	slice := array[1:4]
	fmt.Println(slice, array)

	slice = append(slice, 6, 7, 8) // new underlying array allocated for this slice, the old array will be digested by go's gc
	fmt.Println(slice, array)

	var a []int
	a = append(a, 0)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 2, 3, 4)
	fmt.Println(a, len(a), cap(a))

	for i, v := range slice {
		fmt.Printf("%d %d\n", i, v)
	}

	b := make([][]uint8, 3)
	for row := range b {
		b[row] = make([]uint8, 4)
	}
	fmt.Println(b)
}

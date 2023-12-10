package main

import (
	"fmt"
)

func Index[T comparable](s []T, x T) int {
	for index, val := range s {
		if val == x {
			return index
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 14, 555, 1212}
	fmt.Println(Index(a, 14))

	b := []string{"schmeichel", "degea", "vandersar"}
	fmt.Println(Index(b, "barthez"))
}

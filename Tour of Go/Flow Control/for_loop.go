package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	var sum2 int
	for sum2 < 100 {
		sum2 += 1
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

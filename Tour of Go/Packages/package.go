package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Random variable:", rand.ExpFloat64())
	fmt.Println("Random integer in the range of 0 to 7:", rand.Intn(10))
}

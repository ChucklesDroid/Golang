package main

import (
	"fmt"
	"math"
)

func main() {
	var i rune = 3
	var x rune = 4
	var f float64 = math.Sqrt(float64(i*i + x*x))
	fmt.Println(i, f)
}

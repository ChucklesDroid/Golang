package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10
	var b float64 = 2
	c := b / a
	d := -c

	fmt.Println(math.Ceil(c))
	fmt.Println(math.Floor(c))
	fmt.Println(math.Min(a, b))
	fmt.Println(math.Max(a, b))
	fmt.Println(math.Abs(d))
	fmt.Println(math.Sqrt(b))
	fmt.Println(math.Pow(b, a))

	// complex numbers
	e := complex(3, 4)
	fmt.Println(e)
}

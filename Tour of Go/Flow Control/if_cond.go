package main

import (
	"fmt"
	"math"
)

func sqrt(x int) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(float64(x)))
}

// if with short statement
func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// although v is not accessible here
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

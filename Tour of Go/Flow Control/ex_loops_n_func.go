package main

import (
	"fmt"
	"math"
)

// Newton Raphson's method for Calculating Square Roots
// z -= (z*z - x) / 2z
func Sqrt(x float64) float64 {
	z := float64(1)
	delta := z
	for delta > 0.001 {
		z -= ((z * z) - x) / (2 * z)
		delta = z - math.Sqrt(x)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

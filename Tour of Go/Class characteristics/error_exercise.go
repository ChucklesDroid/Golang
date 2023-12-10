package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	err := float64(e)
	return fmt.Sprintf("cannot Sqrt negative number, %v", err)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	delta := z
	for delta > 0.001 {
		z -= ((z * z) - x) / (2 * z)
		delta = z - math.Sqrt(x)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

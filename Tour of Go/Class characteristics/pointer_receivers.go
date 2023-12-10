package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 { // value receiver
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(multiplier int) { // pointer receiver
	v.X *= float64(multiplier)
	v.Y *= float64(multiplier)
}

func main() {
	a := Vertex{5, 12}
	fmt.Println(a, a.Abs())
	a.Scale(100)
	fmt.Println(a)
}

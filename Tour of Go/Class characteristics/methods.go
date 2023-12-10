package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Feet int

func (a Feet) Square() Feet { // declaring methods on non struct types
	return a * a
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Vertex = Vertex{3, 4}
	fmt.Println(a.Abs())

	b := Feet(2)
	fmt.Println(b.Square())

	c := Myfloat(-2)
	fmt.Println(c.Abs())
}

package main

import (
	"fmt"
)

type Base interface {
	getter() float64
}

type Vertex struct {
	x float64
	y float64
}

func (v *Vertex) getter() float64 {
	return v.x
}

type Myfloat float64

func (f Myfloat) getter() float64 {
	return float64(f)
}

func main() {
	var a Base
	b := Vertex{5, 12}
	a = &b // `&` since its pointer receiver
	fmt.Println(a.getter())

	c := Myfloat(2)
	a = c // `&` since its value receiver
	fmt.Println(a.getter())
}

package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("Vertex with (%v, %v)", v.x, v.y)
}

func main() {
	a := Vertex{5, 12}
	fmt.Println(a)
	fmt.Printf("%s\n", a)
}

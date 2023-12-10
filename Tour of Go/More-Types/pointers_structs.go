package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	b := &a
	b.X = 1e8
	fmt.Println(b.X) // This is effectively (*a).X
	fmt.Println(b.Y)
}

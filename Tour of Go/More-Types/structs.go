package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	var v Vertex = Vertex{2, 3}
	v.x = 2
	fmt.Println(v.x)
	fmt.Println(v)
}

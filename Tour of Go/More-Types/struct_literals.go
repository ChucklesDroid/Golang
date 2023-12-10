package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	var (
		ver1 = Vertex{}     // 0,0 implicit
		ver2 = Vertex{X: 2} // Y = 0 implicit
		ver3 = Vertex{Y: 1} // X = 0 implicit
		ver4 = Vertex{X: 1, Y: 1}
		ver5 = &Vertex{2, 4}
	)
	fmt.Println(ver1, ver2, ver3, ver4, ver5)
}

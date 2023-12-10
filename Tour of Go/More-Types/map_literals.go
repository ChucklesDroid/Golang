package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var m = map[string]int{
	"Hello": 1,
	"World": 2,
	"Manny": 3,
	"Dodo":  4,
}

func main() {
	fmt.Println(m)

	var a = map[string]Vertex{
		"Hello": Vertex{1, 2},
		"World": Vertex{2, 3},
	}

	// The above can be replaced with following
	var b = map[string]Vertex{
		"Hello": {1, 2},
		"World": {2, 3},
	}
	fmt.Println(a, b)
}

package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	dict := make(map[string]Vertex)
	dict["Hello"] = Vertex{3, 4}
	fmt.Println(dict, dict["Hello"])
}

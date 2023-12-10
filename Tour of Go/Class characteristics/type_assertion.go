package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = 42
	v, ok := i.(int)
	fmt.Printf("%v %T %v\n", v, i, ok)

	a := i.(float32) // causes panic
	fmt.Printf("%v %T\n", a, i)
}

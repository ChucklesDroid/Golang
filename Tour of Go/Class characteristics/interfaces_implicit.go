package main

import (
	"fmt"
)

type I interface {
	M()
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I = F(2)
	i.M()
}

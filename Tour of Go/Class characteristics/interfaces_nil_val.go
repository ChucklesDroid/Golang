package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	s string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

type A struct {
	message string
}

func (a A) M() {
	fmt.Println(a.message)
}

func analyzer(i I) {
	fmt.Printf("%v %T\n", i, i)
}

func main() {
	var i I  // declaring interface
	var t *T // declaring struct object whose zero value is nil
	i = t    // interface with nil value
	i.M()

	var a A
	analyzer(a)
}

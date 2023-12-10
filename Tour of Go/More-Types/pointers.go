package main

import (
	"fmt"
)

func main() {
	var a int = 8
	p := &a
	fmt.Printf("*p: %d\n", *p)
	*p = 19
	fmt.Println(*p)
}

package main

import (
	"fmt"

	"examples.com/integers"
	// usual syntax:= host/usrororganization/project/(dirs)/package
)

func main() {

	var integer int = 10
	fmt.Println(integer)

	// Block Scope
	// This is called block scope in which elements are not accessible outside the scope
	// {
	// 	var integer2 int
	// }
	// fmt.println(integer2)

	//making use of exported variable:
	fmt.Println(integers.Three)
}

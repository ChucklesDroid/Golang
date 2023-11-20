package main

import "fmt"

var a, b int // Declared in the package's scope

func main() {
	var c int = 1   // Declared and Initialised in the function scope
	var d, e = 3, 4 // If an initialiser is present type can be omitted
	fmt.Println(3)
	fmt.Println(c)
	fmt.Println(d, e) //Multiple Initialisers in a single line

	// Short Declaration can be used instead of var (Can be used inside a function)
	f := 9
	fmt.Println(f)
}

package main

import "fmt"

func main() {
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius

	// c = f //produces error since c and f are of diff types

	// To convert it effectively into celsius use the following code:
	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}

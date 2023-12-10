package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Creating slices
	var b []int = primes[3:5] // Last element in range is not included
	fmt.Println(b)

	var names [4]string = [4]string{
		"Aakarsh",
		"Dimes",
		"Manny",
		"Jerry",
	}

	var d []string = names[2:4]
	var e []string = names[1:3]
	fmt.Println(d, e)

	// Slices only act as reference
	d[0] = "XXXX"
	fmt.Println(d, e)
}

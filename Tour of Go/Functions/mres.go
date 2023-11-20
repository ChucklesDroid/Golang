package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("hello", "world") // We make use of ":=" to declare and initialise a variable
	fmt.Println(a, b)
}

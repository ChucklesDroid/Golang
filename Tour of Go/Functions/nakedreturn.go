package main

import "fmt"

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func split(x int) (int, int) {
	a := x % 10
	b := x - a
	return a, b
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(split(78))
}

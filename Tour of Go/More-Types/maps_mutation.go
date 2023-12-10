package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{
		"Hello": 1,
	}
	fmt.Println(m)

	// Insert or update an element
	m["World"] = 2
	fmt.Println(m)

	m["Hello"] = 2
	fmt.Println(m)

	// Retrieve an element
	a := m["Hello"]
	fmt.Println(a)

	// Delete an element
	delete(m, "Hello")
	fmt.Println(m)

	// 2 value assignment
	// When key is not present in map:
	// a) val = 0
	// b) ok = false

	// When key is present in map: // val, ok = m["World"]
	// a) val = 2
	// b) ok = true
	val, ok := m["Hello"]
	fmt.Println(val, ok)

	b := make(map[string]int)
	b["Hello"] = 2
	fmt.Println(b)
}

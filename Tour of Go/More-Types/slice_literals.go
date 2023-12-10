package main

import (
	"fmt"
)

func main() {
	var (
		a []bool = []bool{true, false, false, true, true}
		b []int  = []int{1, 2, 3, 5, 9}

		// s []struct {
		// 	X string
		// 	Y bool
		// } = []struct {
		// 	X string
		// 	Y bool
		// }{
		// 	{"Aakarsh", true},
		// 	{"Manny", false},
		// 	{"Ramen", true},
		// }
	)
	s := []struct {
		X string
		Y bool
	}{
		{"Aakarsh", true},
		{"Manny", false},
		{"Ramen", true},
	}

	fmt.Println(a, b, s)
}

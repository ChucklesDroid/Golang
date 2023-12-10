package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Not string nor int")
	}
}

func main() {
	var i interface{}
	i = 43
	do(i)
}

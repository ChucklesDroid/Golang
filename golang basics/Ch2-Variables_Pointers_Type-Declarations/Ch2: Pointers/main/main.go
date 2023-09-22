package main

import "fmt"

func main() {
	var integer int
	var p *int = &integer

	fmt.Println(integer, p)
}

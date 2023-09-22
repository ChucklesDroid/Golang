package main

import "fmt"

func main() {
	var x = 1
	fmt.Println(x)
	var x32 int32 = 2
	fmt.Println(x32)
	var x64 int64
	fmt.Println(x64)

	fmt.Printf("%T, %T, %T\n", x, x32, x64)

	// Conversion
	// x = x32 // Error: cannot use x32 (type int32) as type int in assignment
	x = int(x32)
	fmt.Println(x)

	y := int64(x32)
	fmt.Println(y)

	// Unsigned int
	var z uint = 10
	// x := -1 // Error: constant -1 overflows uint
	fmt.Println(z)

	fmt.Println(x + int(y))
	fmt.Println(x % int(y))
	fmt.Println(x == int(y))
	fmt.Println(x < int(y))

	//Floating number points
	var a = 3.15
	fmt.Printf("%T\n", a)
	var b float32 = 2.13
	fmt.Printf("%T\n", b)
	var c float64 = 2.77
	fmt.Printf("%T\n", c)

	//Exponential notation
	d := 13e12
	fmt.Println(d)

	//Floating point conversion
	e := float32(c)
	fmt.Println(e)

	f := int(d)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}

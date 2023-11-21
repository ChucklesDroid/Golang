package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
)

func main() {
	// var boolean bool = false
	// var a int = -2  // integer 32 bit
	// var b uint = 2  // unsigned integer 32 bit
	// var c byte = 8  // 8 bit integer
	// var d rune = 32 // 32 bit integer
	// var e *int = &a

	// var f float32 = 3.1
	// var g float64 = 32.1

	// var h complex64 = 1 + 2i

	// We can also do the following
	var (
		boolean bool      = true
		a       int       = -2
		b       uint      = 2
		c       byte      = 8
		d       rune      = 32
		e       *int      = &a
		f       float32   = 3.2
		g       float64   = 32.2
		h       complex64 = 1 + 2i
		i       bool      = cmplx.IsNaN(complex128(h))
	)

	fmt.Println(boolean)
	fmt.Print(a, b, c, d, e, f, g, h, i, "\n")
	fmt.Println(reflect.TypeOf(e)) // To print variable type
}

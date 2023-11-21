package main

import (
	"fmt"
)

// const Big = 1 << 100
// const Small = Big >> 99

// You can replace the above code with the following
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	// The following statement results in error since it overflows
	fmt.Println(int(Big)) // As you can see it overflows however we can declare cosnstant variables with such values without overflowing
}

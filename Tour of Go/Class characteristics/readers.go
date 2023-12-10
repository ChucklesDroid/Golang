package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	readers := strings.NewReader("Hello, Reader!") // creating the underlying data source, from which to read from

	p := make([]byte, 8)
	for {
		num, err := readers.Read(p)
		if err == io.EOF {
			fmt.Println("End of File reached !")
			break
		} else {
			fmt.Printf("%v bytes were read and are as follows: %v\n", num, string(p))
		}
	}
}

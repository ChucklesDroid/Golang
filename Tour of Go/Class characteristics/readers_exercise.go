package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	read, err := rot.r.Read(p)
	if err != nil && err != io.EOF {
		fmt.Printf("Not Working, %v", err)
		return 0, err
	} else if err == io.EOF {
		return 0, io.EOF
	} else {
		for i := 0; i < read; i++ {
			if p[i] >= byte('A') && p[i] < byte('N') || p[i] >= byte('a') && p[i] < byte('n') {
				p[i] += byte(13)
			} else if p[i] >= byte('N') && p[i] <= byte('Z') || p[i] >= byte('n') && p[i] <= byte('z') {
				p[i] -= byte(13)
			}
		}
		return read, nil
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

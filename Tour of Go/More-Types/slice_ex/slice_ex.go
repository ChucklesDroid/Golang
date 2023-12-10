package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for row := range slice {
		slice[row] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch slice[i][j] % 2 {
			case 0:
				slice[i][j] = uint8(i * j / 2)
			default:
				slice[i][j] = uint8(i ^ j)
			}
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}

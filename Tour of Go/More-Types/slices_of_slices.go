package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	printSlice(board)
}

func printSlice(s [][]string) {
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Print(s[i][j])
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(strings.Join(s[i], " "))
	}
}

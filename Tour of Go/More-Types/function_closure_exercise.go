package main

import "fmt"

func fibonacci() func() int {
	var (
		first  int = 0
		second int = 0
	)
	return func() int {
		if first+second == 0 {
			second = 1
			return first
		} else if first+second == 1 || (first <= 1 && second == 1) {
			first++
			return second
		} else {
			if first > second {
				first = second
				second = first + second
			} else {
				temp := first
				first = second
				second = temp + second
			}
			return second
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

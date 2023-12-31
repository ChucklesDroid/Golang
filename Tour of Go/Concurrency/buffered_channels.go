package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

package main

import (
	"fmt"
)

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	channel := make(chan int)
	go sum(2, 3, channel)
	go sum(3, 5, channel)
	x, y := <-channel, <-channel
	// z := <-channel // not allowed no sender
	fmt.Println(x, y)
	go func() {
		channel <- 5
	}()
	fmt.Println(<-channel)
}

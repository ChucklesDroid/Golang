package main

import (
	"fmt"
)

func Init(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // To use it in conjunction with range its essential to use close() to close the channel
}

func main() {
	channel := make(chan int, 10)
	Init(channel)
	for val := range channel {
		fmt.Println(val)
	}
	channel2 := make(chan int)
	close(channel2)
	val, ok := <-channel2
	fmt.Println(val, ok) // value returned from a close channel is zero value of that channel
}

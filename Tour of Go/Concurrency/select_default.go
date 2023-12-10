package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom.")
			return // by using a break here it will only break out of the current iteration of the select statement
		default: // Therefore to break out use return
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

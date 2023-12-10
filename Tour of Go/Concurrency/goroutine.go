package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	go say("Hello")
	say("World")
}

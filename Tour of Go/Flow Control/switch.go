package main

import (
	"fmt"
	"time"
)

// runtime.GOOS constant returns the os type. It can be read as Go's os
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// fmt.Print("Go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	// FREEBSD, UNIX, OPENBSD, WINDOWS
	// 	fmt.Printf("%s\n", os)
	// }
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day after Tomorrow")
	default:
		fmt.Println("Too far away")
	}
}

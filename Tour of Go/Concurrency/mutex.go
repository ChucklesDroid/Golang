package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	lock sync.Mutex
	v    map[string]int
}

func (s *SafeCounter) Inc(k string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.v[k]++
}

func (s *SafeCounter) Value(k string) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.v[k]
}

func main() {
	a := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go a.Inc("Key")
	}
	time.Sleep(time.Second) // Since we are not making use of channels to await the completion of previous go routine
	fmt.Println(a.Value("Key"))
}

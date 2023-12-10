package main

import (
	"fmt"
)

type List[T any] struct {
	next  *List[T]
	value T
}

func (L *List[T]) Append(head *List[T]) {
	for L.next != nil {
		L = L.next
	}
	L.next = head
}

func main() {
	a := List[int]{value: 2}
	b := List[int]{value: 3}
	c := List[int]{value: 4}
	a.Append(&b)
	a.Append(&c)
	fmt.Println(a, a.next, b.next)
}

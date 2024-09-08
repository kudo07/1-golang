package main

import "fmt"

type PersonP struct {
	name string
	age  uint
}

func (p PersonP) Equal(p2 PersonP) bool {
	return p.age == p2.age && p.name == p2.name
}

func people() {
	slice := []PersonP{{"jOe", 10}, {"jOe", 10}}
	fmt.Println(slice[0].Equal(slice[1]))
}

// true

package main

import "fmt"

func completePointer() {
	x := 0
	y := 1
	s := []*int{&x, &y}
	// in the slice the valriabe addresses are stored automatically by implicitly get the pointer of that variable location hesaddress
	// these are the slice of the pointers
	fmt.Println(s)
	// [0xc00000a160 0xc00000a168]
}

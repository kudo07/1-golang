package main

import "fmt"

func changeSlice(x []int) {
	x[0] = 1
}

func slicee() {
	s := []int{0, 0, 0}
	changeSlice(s)
	fmt.Println(s)
	// 0xc00000a138
}

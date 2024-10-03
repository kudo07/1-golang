package main

import "fmt"

func pointerSlice() {
	x := []string{"hello", "world"}

	y := &x
	// y is storing pointer to x
	z := &y
	// and z is storing pointer to gives me the pointing to x
	// defreference of z g

	// pointer to the pointer of the y and y has the pointer of the x
	fmt.Println(x, y, z)
	// [hello world] &[hello world] 0xc00005c058
	fmt.Println(x, *y, *z)
	// [hello world] [hello world] &[hello world]

	fmt.Println(x, *y, **z)
	// [hello world] [hello world] [hello world]
	fmt.Printf("%T", z)
	// **[]string
	// 22:36
}

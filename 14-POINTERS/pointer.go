package main

import "fmt"

func pointer() {
	x := 0
	y := &x
	fmt.Println("1", y, &x)
	// 0xc00000a138 0xc00000a138
	// modify x in the pointer of x which stored in y
	fmt.Printf("%T", y)
	*y = 10
	println()
	// *y is the pointer which store the address of x as &x in the varibale called y so
	// y sored in different location with the value of address of x
	fmt.Println(x, y)
	// 1 0xc00008c0f8 0xc00008c0f8
	// *int
	// 10 0xc00008c0f8
	// so  y has act the intermediatory to change the value of other variable and display the change in the x variable
	a := 2
	b := &a
	fmt.Println(a, *b)
	// 2 2

	// because derefernce of b which is pointing to the value of the a give us the value of 2
	// *INT MEANS pointer to an int type
	// *y before a pointer means defereence the pointer for me
}

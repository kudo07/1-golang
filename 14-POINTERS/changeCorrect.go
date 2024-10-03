package main

import "fmt"

func changeddd(x *int) {
	// the type is pointer
	// as a oaram accepting the pointer to the int value
	*x = 4
}

func changeCorrect() {
	y := 2
	changeddd(&y) // passs the addres of y
	fmt.Println("correctchanged", y)
}

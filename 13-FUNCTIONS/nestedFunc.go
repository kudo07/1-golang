package main

import "fmt"

// accept type in the braces and return the int in the outside
func returnFunc(x int) func(int) int {
	return func(y int) int { return x + y }
}

func nestedFunc() {
	fmt.Println("hi")
	// access the function as the parameter type  as the function
	fn := returnFunc(100)
	value := fn(50)
	fmt.Println(value)
}

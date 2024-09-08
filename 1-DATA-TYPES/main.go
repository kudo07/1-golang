package main

import "fmt"

func main() {
	var x string
	x = 12
	// error
	var y int
	var z float32
	fmt.Println(x, y, z)
	i := int8(128)
	fmt.Println(i)
}

// statistically type golang explicit type defined

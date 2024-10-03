package main

import "fmt"

func main() {
	x := 2
	y := x
	x = 3
	fmt.Println(x, y)
	change()
	changeCorrect()
	slicee()
	pointer()
	pointerSlice()
	structPointer()
	pointerStruct()
	pointerStructTwo()
	complexPointer()
	completePointer()
}

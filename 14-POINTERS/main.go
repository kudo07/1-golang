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

// go run main.go change.go changeCorrect.go slicee.go pointer.go pointerSlice.go structPointer.go pointerStruct.go pointerStructTwo.go complexPointer.go completePointer.go

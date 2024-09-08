package main

import "fmt"

func slices() {
	fmt.Println("hello")
	// pointer
	// length
	// capacity is the underlined array
	arr := [5]int{1, 2, 3, 4, 5}
	sliceArr := arr[:]
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
	sliceArr = arr[1:4]
	sliceArr[1] = 20020

	sliceArr[2] = 9898
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
	// LEN IS THE LENGTH OF THE SLICE NEW ARRAY
	//
}

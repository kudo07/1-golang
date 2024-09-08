package main

import "fmt"

func rangee() {
	arr := []string{}
	for i := 0; i < 10; i++ {
		arr = append(arr, "hello")
		fmt.Println(arr, len(arr), cap(arr))
	}
	arr1 := make([][]int, 5)
	var x string     // declaration or initialisation
	x = "hello word" // assignment
	// all in one
	e := "hello wold"
	fmt.Println(e, x, arr1)
	raa := make([][]int, 5, 10)
	// 5 is length and 10 is the capacity capacity increases due to increase in the number of element

	fmt.Println(raa, len(arr), cap(arr))
	drr := []bool{}
	// this is the slice
	drr2 := []bool{true, false, false}
	// this is the slice
	drr = append(drr, drr2...)
	//
	fmt.Println(drr)
	rrr := []bool{true, false, false}
	for i, val := range rrr {
		fmt.Println(i, val)
	}
	change(rrr)
	fmt.Println(rrr)
	// not change
	rrr1 := [3]bool{true, false, false}
	change1(rrr1)

	fmt.Println(rrr1)
	// [false false false]
	// it really change because we are actually change the underlined array
	// }looking at 0=false

	// slice is the view of the underlined array
	// increasing the size of the array will lead to increasing the size of the underlined aray
	// reason to store the value in the array
	// the slice are meant to abstract on top
	// flexible allow us to accept as a parameter a slice to store any number of elements inside it
}

func test(x [2]int) {
	// do something
	// inflexible only store 2 element inside the array
}

func test2(x []int) {
	// flexible store any number of elements inside it
}

// we did actually change the element this time
// we change the array in the slice [] actually modifying the underlined array or able to mutate the array inside the function
// we change the slice whathappens is  we modifies the underlined array that we able to mutate the array inside the finction
// when we pass the slice u can directly change the array inside the function but when you pass the array  u cannot pass as the copy
func change(x []bool) {
	x[0] = false
}

// no change were  made because once we pass the array to the function we are passing the copy of the array which is "x"
// we are no the passing the rrr1 we are passing the copy of the rrr1 array which doent change the rrr1
func change1(x [3]bool) {
	x[0] = true
}

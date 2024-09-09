package main

import "fmt"

// return function from other function and nested function
func test(a int, b int) (x int, y int) {
	x, y = a, b
	return
	// since x has a value and y has a value it automatically return both x and y in the order which has above
	// even dont assign a value because implecitly it creeate the variabel x and y has the value with int type
}

// unlimited nums
func sumInts(nums ...int) int {
	fmt.Println(nums)
	return 1
}

// variadic functions
func sumIntsOne(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("hi")
	a, b := test(1, 2)
	fmt.Println(a, b)
	// 1 2
	sumInts(1, 2, 3, 4, 4, 4)
	// slice of number
	// [1 2 3 4 4 4]
	sumOne := sumIntsOne(1, 2, 2, 3, 342, 432)
	fmt.Println(sumOne)
	// lets break the slice
	nums := []int{1, 2, 32, 43, 4234}
	sumTwo := sumIntsOne(nums...)
	fmt.Println(sumTwo)
	// 4312
	// work in the same way
	nestedFunc()
	// it makes the closure
	closure()
	// parameter in the function
	parametersFunc()
}

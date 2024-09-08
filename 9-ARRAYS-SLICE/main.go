package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var arr [2]int
	arr[0] = 100
	fmt.Println(len(arr))
	// 2
	fmt.Println(arr)
	//[100,2]
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{12, 3, 4, 5, 5, 5, 55}
	fmt.Println(arr1)

	fmt.Println(arr2)
	arr3 := [2][2]string{{"hello", "world"}, {"code", "golang"}}
	fmt.Println(arr3)

	fmt.Printf("%T", arr3)
	// not [][]string but rather define the number of elements
	// looping in the arrays
	for o, nestedArr := range arr3 {
		fmt.Println("")
		fmt.Println(o, nestedArr)
		for j, num := range nestedArr {
			fmt.Println("next loop")
			fmt.Println(j, num)
		}
	}

	arr4 := [2][2]string{{"change", "world"}, {"new", "golang"}}
	fmt.Println("arr4", arr4)
	slices() // call the function from the slices.go
	rangee()
}

func repair(x [2][2]string) {
	// i pass the array which is copy of arr4 in x
	x[0][0] = "repair"
	// the change in x not gonna make change in the arr4
}

package main

import "fmt"

func main() {
	fmt.Println("here we go")
	var mp map[string]int
	fmt.Println(mp)
	// map[]
	// 2
	mp1 := map[string]int{"h": 1, "a": 3}
	fmt.Println(mp1)
	// map[a:3 h:1]
	// 3
	mp2 := make(map[int]rune)
	mp2[5] = 2
	fmt.Println(mp2)
	// map[5:2]
	delete(mp2, 5)
	fmt.Println(mp2)

	// map[]
	// blunt doesnt give error
	delete(mp2, 3)
	// same ma2 but didnt give an error
	mp2[10] = 3
	value, ok := mp2[10]
	fmt.Println(value, ok)
	// 3 true
	// 4
	mp3 := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range mp3 {
		fmt.Println(key, value)
	}

	// 5
	n := 100
	divisibleMap := make(map[int]uint)
	for num := 1; num <= n; num++ {
		for d := 1; d <= 5; d++ {
			if num%d == 0 {
				divisibleMap[d]++

				// fmt.Println(divisibleMap, num, d)
			}
		}
	}
	// like every number is divisible by 100 si 1:100 50% of 100 make divisible by 2 and so on
	fmt.Println(divisibleMap)
}

// rune is an alias for int32

package main

import "fmt"

func changedd(x int) {
	x = 4
}

func change() {
	y := 2
	changedd(y)
	fmt.Println(y)
	// y will still be 2, not 4
	// the function change(x int ) takes an integer x as an argument but only
	// modifies its local copy
	// when  y:=2 is passed to the change function, the function change the local copy but the original value of y remains unchanged
	// therefore when you print y it will still output 2
}

// x int is the copy whatever pass from the call

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// the closure is bound to this variable sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	// both are pointing thier differ adder repectively
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// a closur eis a function value that reference variable from outside its body the function
// may crosss.and assign to the referenced variables in this sense is "bound" to the variable
// adder function returns a closure/ Each closure is bound to its own sum variable
// so the pos and neg has the return function value from the adder function which has anonymous function iside it
// the concept is each time we call now after the initialistion of pos and neg
// we get the value in pos it has sam but neg has 2*(-i)
// means pos and neg pointing the same value as it point the valu from the valoue respectively

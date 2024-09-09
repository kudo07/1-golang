// functions accept the parameter as the function
package main

import "fmt"

// the arg string which is parameter is the argument to the callable function
func callFunc(callable func(string) string, arg string) {
	result := callable(arg)
	// the callable will call the psasing function and pass the arg sting to ot whcih made str+ hello
	// later the str is "world " which pass this callable transfer to it and make it world hello
	fmt.Println(result)
}

func parametersFunc() {
	// creating the function
	myFunc := func(str string) string {
		return str + "hello!!!"
	}
	callFunc(myFunc, "world")
}

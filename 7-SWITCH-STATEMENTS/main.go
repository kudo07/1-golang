package main

import "fmt"
func main(){
	// a:=1
	// switch a{
	switch a:=0; {
	case a<1:
		fmt.Println("this is one ")
		fallthrough
		// itt automatically executes what inside the next case
	case a<2:
		fmt.Println("this is two")
	default:
		fmt.Println("this is default")
	}
	fmt.Println("Hello")
	// check the multiple case in one statement
	
}
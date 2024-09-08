package main

import "fmt"
func main(){
	var x int=1
	// explicitly you cannot change the type of the varible later in the program

	y:=2
	// essentially the same thing this is implicit way
	// guess whats the type of the "y" be   , [preffered]
	// y:="fwerfewr" 
	// implecit guessing the type of the variable
	//  :=  compiler guess whats the type of the value
	// we will get the type int
	//  but y could be int float uint
	// .2 ways either you go back and declare the variable or 
	a:=uint(2)
	z:='e'
	s:=byte('3')
	fmt.Printf("%T",z)
	// int32
	fmt.Printf("%T",s)
	// uint8

	fmt.Println(x,y,a,z)
}
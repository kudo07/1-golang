package main

import "fmt"
func main(){
	x:=1
	fmt.Printf("%T %T",x,x)
	// seperate %t for the lineup x
	// int int 
	// type
	// printf is for the formatting 
	y:=true 
	z:=12
	fmt.Printf("%v hello %T",x,y)
	fmt.Printf("%v hello %T %b",x,y,z)
	// binary representation   
	u:=12.2822432323
	fmt.Printf("%e %.2f",z,u)
	// 12.28
	// 2 point precision
	fmt.Printf("%e %%%10.2f",z,u)
	// width of 10

}

func defineOutput(x string) (string , int){
	return x,1
}
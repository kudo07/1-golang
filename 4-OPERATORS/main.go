// to add the two variables the two variable must have the same tyoes
package main

import (
	"fmt"
	"strconv"
)
func main(){
	// a:=5
	// b:=int8(2)
	// // mismatch types
	// c:=a/int(b)
	// // now working 
	// fmt.Println(c)
	// SAME TYES STATISTICVALLY TYPED LANG
	// conver string into number
	// str:="123hello"

	str1:="123"
	// a,err:=strconv.Atoi(str1)
	// fmt.Printf("%T %T",a,err)
	b,err1 := strconv.ParseInt(str1,10,64)
	fmt.Println(b,err1)
}
package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func Test() string {
	return ""
}

// this will automatically export the function outside the main package
// capital letter structs and functions can be accessed out side the package main
// note to access the attributes they should name as the capital other wise they dont export it and use indirectly as the dot operator
func (p Person) GetName() string {
	return p.name
}

func (p Person) SetName(newName string) {
	p.name = newName
	fmt.Println(p)
}

// p acted as Person
func main() {
	p1 := Person{age: 21, name: "Tim"}
	fmt.Println(p1)
	// 1 var
	var p2 Person
	p2.name = "Tim"
	fmt.Println(p2)
	// eported
	p3 := Person{"time", 123}
	fmt.Println(p3.GetName())
	//
	p3.SetName("ioui")
	// i am passsing as p is the copy of the p3 struct
	// so p is the same info but the ciopoy of the p3 struct
	//it will change that p but not gonna change the p3
	//
	fmt.Println(p3.GetName())
	// strange we get the "time" here but got upper {ioui 123}
	// not .working here
	clothingInherit()
	people()
}

// go run main.go clothingInherit.go

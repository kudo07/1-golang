package main

import "fmt"

// interface
type Name interface {
	GetName() string
	// GetLastName() string
	// gives errro se we have to implement it into both structs
}

// struct
// person
type Person struct {
	firstName string
	lastName  string
}

// employee
type Employee struct {
	name string
}

// methods
// this method needs to be implemented to make the Person struct run corrextly
// person
func (p Person) GetName() string { return p.firstName + " " + p.lastName }

func (p Person) SayHi() {
	fmt.Println("Hi!")
}

//

// employee
func (e Employee) GetName() string {
	return e.name
}

func PrintName(obj Name) {
	fmt.Println(obj.GetName())
	// obj with the interface Name
	// dont care what type of obj is either Person or Employee
	// multiple objects or multiple structs at the same time
}

// ṃain
func main() {
	// the type of the variable name is Name which equals to the Person struct
	var name Name = Person{"Tim", "fwef"}
	// we say this person struct struct as the type name so we cant access name.firstName it gives error
	// since it implement as the interface
	// to hide certain aspects and complexity
	// shields the person struct
	employeeNamee := Employee{"emplyee234324"}
	fmt.Println(employeeNamee)
	fmt.Println(name.GetName())
	// Tim fwef
	// i want to make the slice of the name type
	//
	// Employee and Person are bothe the structs
	// Employee and Person has seperately two methods which return the same thing with same Name
	// "Name" GetName is the name of both the methods and "Name " is the name of the Interface
	//both of the structs inplemented the Name interface
	// so anything Name interface which has slice of names is valid to be in the slice
	//
	names := []Name{Employee{"Tim"}, Person{"Lin", "kin"}}
	fmt.Println(names)
	// to loop in the names and extract the GetName
	for _, name := range names {
		fmt.Println(name.GetName())
	}
	// tim
	// lin kin
	// .view different structs in the same way
	// emplpoyee and person may be have different methods and work in the different manner
	// viewing as the same type of object which is "Name" and allow me to accces any of the methods inside the "Name " interface
	// if i implement or define the different methods in interface then we have to define the methods as in the struct of employee adn person
	// but that is not mean that these structs not have their specif methods
}

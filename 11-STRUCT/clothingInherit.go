package main

import "fmt"

type PersonO struct {
	name     string
	age      uint
	clothing Clothing
	// this is called embeding structs
}
type Clothing struct {
	shoeSize   uint
	shirtColor string
}

func (c Clothing) GetShoeSize() uint {
	return c.shoeSize
}

func (p PersonO) GetShoeSize() uint {
	return p.clothing.GetShoeSize()
}

// t still works
func clothingInherit() {
	p4 := PersonO{"TIMTI", 2323, Clothing{12, "orange"}}
	shoeSize := p4.GetShoeSize()

	// shoeSize := p4.clothing.GetShoeSize()
	fmt.Println(p4)
	fmt.Println(shoeSize)
	// 12
}

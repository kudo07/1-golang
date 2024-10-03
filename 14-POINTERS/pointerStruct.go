package main

import "fmt"

type Book1 struct {
	title string
	id    int
}

func changeTitle1(book1 *Book1, title string) {
	// here i derefernce it by getting the address of that value to make impact on the original parent
	// book1 is the pointer here whenever we accessing the field of the struct ist automatically derefernce it by the dot notation

	book1.title = title
	// we dont write like this *book1.title as when we pass the astric in the params of strct so automatically
	// it dereference the field like this case in title
	// i dont have to manually derefernce it
	// however manually like this
	// (*book1).title = title
}

func pointerStruct() {
	book1 := Book1{"time is great!", 1}
	changeTitle1(&book1, "new title")
	// here i pass the reference and
	// {new title 1}
	fmt.Println(book1)
}

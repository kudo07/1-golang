package main

import "fmt"

type Took1 struct {
	title string
	id    int
}

// write a methods by using the pointer to a struct
func (b *Took1) titleChange1(title string) {
	// accepting a pointer to a Took actually modified the underlined Took
	b.title = title
	// automatically dereferencey
	// if we remove the pointer then it actualy become the copy and didnt effect on original Toook
}

func complexPointer() {
	took1 := []*Took1{{"a", 1}}
	// slicise that are stores the pointers of the book
	// write the bunch of books inside it
	// implicitly giving the pointer
	// i dont have to manually do []*Took1{&Took1{"a",1}}
	// this is also valid code this is explicitly saying that i want to get the pointer to the took1 that i just created
	// get the pointr of the book that i just created this is explicity by &Took1{'a',1} but we just implicitly saying by ermoving the pointer
	// leave it implicitly give the pointer of the book create objext because define the type of the slice of the Took1
	// slice of a pointer which stores a book
	fmt.Println(took1)
	// {title new 1}
}

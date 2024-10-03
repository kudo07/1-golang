package main

import "fmt"

type Took struct {
	title string
	id    int
}

// write a methods by using the pointer to a struct
func (b *Took) titleChange(title string) {
	// accepting a pointer to a Took actually modified the underlined Took
	b.title = title
	// automatically dereferencey
	// if we remove the pointer then it actualy become the copy and didnt effect on original Toook
}

func pointerStructTwo() {
	took := Took{"i am great!", 1}
	took.titleChange("title new")
	fmt.Println(took)
	// {title new 1}
}

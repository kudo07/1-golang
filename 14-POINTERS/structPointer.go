package main

import "fmt"

type Book struct {
	title string
	id    int
}

func changeTitle(book Book, title string) {
	// book Book is the copy of book struct
	// the book is not here which is passed below it is copy
	// changing the copy in it
	// it is out of the scope so the 'new title' is not changing
	book.title = title
}

func structPointer() {
	book := Book{"I ma great!", 1}
	changeTitle(book, "new title")
	fmt.Println(book)
	// []string{I ma great! 1}
}

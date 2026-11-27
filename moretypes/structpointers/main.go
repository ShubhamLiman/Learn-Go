package main

import "fmt"

func main() {
	book := Book{"Go Basics", 100}
	addPagesCopy(book)
	fmt.Println(book.Pages)
	addPagesPointer(&book)
	fmt.Println(book.Pages)
}

func addPagesCopy(b Book) {
	b.Pages += 10
}

func addPagesPointer(b *Book) {
	b.Pages += 10
}

type Book struct {
	Title string
	Pages int
}

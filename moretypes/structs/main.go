package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func main() {
	b := Book{"Go Basics", 120}
	fmt.Println(b)
	b.Pages = 150
	fmt.Println(b.Title, b.Pages)

	var empty Book
	fmt.Println(empty)
}

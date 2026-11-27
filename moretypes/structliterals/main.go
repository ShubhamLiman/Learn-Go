package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func main() {
	b1 := Book{Title: "Go Basics"}
	fmt.Println(b1)
	b2 := Book{Pages: 300, Title: "Rust Book"}
	fmt.Println(b2)
	b3 := &Book{Title: "Pointers", Pages: 50}
	b3.Pages = 75
	fmt.Println(b3)

}

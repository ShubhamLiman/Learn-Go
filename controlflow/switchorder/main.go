package main

import "fmt"

func main() {
	x := 7
	switch x {
	case check(1):
		fmt.Println("one")
	case check(2):
		fmt.Println("two")
	case check(3):
		fmt.Println("three")
	default:
		fmt.Println("no match")
	}
}

func check(n int) int {
	fmt.Println("checking", n)
	return n
}

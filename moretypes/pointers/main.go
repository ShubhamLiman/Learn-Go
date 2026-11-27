package main

import "fmt"

func main() {
	score := 50
	p := &score
	fmt.Println(p)
	fmt.Println(*p)
	*p = *p * 2
	fmt.Println(score)
	var empty *int
	fmt.Println(empty)
}

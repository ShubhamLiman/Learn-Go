package main

import "fmt"

func main() {
	scores := [5]int{10, 20, 30, 40, 50}
	backup := scores
	first := scores[0:3]
	last := scores[2:5]
	fmt.Println(first, last)
	first[2] = 99
	fmt.Println(first, last)
	fmt.Println(scores)
	fmt.Println(backup)

}

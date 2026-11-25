package main

import "fmt"

func main() {
	n := 30
	if half := n / 2; half > 10 {
		fmt.Println("big half", half)
	} else {
		fmt.Println("small half", half)
	}

}

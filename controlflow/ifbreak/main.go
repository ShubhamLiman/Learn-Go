package main

import "fmt"

func main() {
	n := 1
	steps := 0

	for {
		n = n * 2
		steps += 1
		if n > 1000 {
			break
		}

	}
	fmt.Println(n)
	fmt.Println(steps)
}

package main

import "fmt"

func main() {
	fmt.Println("main start")
	countUp()
	fmt.Println("main end")
}

func countUp() {
	count := 0
	defer fmt.Println("deferred count:", count)
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("count:", count)
}

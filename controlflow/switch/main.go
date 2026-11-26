package main

import "fmt"

func main() {
	light := "yellow"

	switch light {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Slow down")
	case "green":
		fmt.Println("go")
	default:
		fmt.Println("broken light")

	}
	fmt.Println("done")
}

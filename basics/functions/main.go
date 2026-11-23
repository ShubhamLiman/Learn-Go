package main

import "fmt"

func rectangleArea(width, height int) int {
	return width * height
}

func minutesAndSeconds(totalSeconds int) (int, int) {

	return totalSeconds / 60, totalSeconds % 60

}

func main() {
	fmt.Println("Area of rectangle", rectangleArea(3, 4))
	minutes, seconds := minutesAndSeconds(192)
	fmt.Println("minutes =", minutes, "seconds =", seconds)
}

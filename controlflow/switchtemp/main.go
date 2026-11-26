package main

import "fmt"

func main() {
	fmt.Println(temperatureLabel(-5))
	fmt.Println(temperatureLabel(0))
	fmt.Println(temperatureLabel(14))
	fmt.Println(temperatureLabel(15))
	fmt.Println(temperatureLabel(25))
}
func temperatureLabel(temp int) string {
	switch {
	case temp < 0:
		return "freezing"
	case temp < 15:
		return "cold"
	case temp < 25:
		return "mild"
	default:
		return "hot"
	}
}

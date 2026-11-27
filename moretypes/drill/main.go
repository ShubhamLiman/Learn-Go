package main

import "fmt"

func main() {
	c := Counter{Name: "clicks"}
	increment(&c)
	increment(&c)
	increment(&c)
	fmt.Println(c)
	d := c
	increment(&d)
	fmt.Println(c.Count, d.Count)

}

func increment(count *Counter) {
	count.Count += 1
}

type Counter struct {
	Name  string
	Count int
}

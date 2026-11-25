package main

import "fmt"

var c, python, java bool
var m, n int = 1, 2
var k = 3

func main() {
	var i int
	fmt.Println(i, c, python, java)

	x := 1
	x, z := 3, 4
	fmt.Println(x, z)

	var isGo, isRust, answer = true, false, "no!"

	fmt.Println(m, n, isGo, isRust, answer)
	fmt.Println(k)

	var a int = 5
	var b int64 = 5
	fmt.Println(a + int(b))

	p := 3.0
	q := 7
	s := b
	g := 'G'
	fmt.Printf("%T %T %T %T\n", p, q, s, g)
	fmt.Println("go", 'G')
}

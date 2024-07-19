package main

import "fmt"

func main() {

	fmt.Println("Declare variable")
	var x int = 10
	var y = 20
	z := 30
	fmt.Println(x, y, z)

	var x2, y2 int = 10, 20
	fmt.Println(x2, y2)

	var (
		x3   int
		y3       = 20
		z3   int = 30
		d, e     = 40, "hello"
		f, g string
	)
	fmt.Println(x3, y3, z3, d, e, f, g)
}

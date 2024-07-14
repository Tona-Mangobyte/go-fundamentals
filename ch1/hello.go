package main

import "fmt"

func main() {

	fmt.Println("Hello, world!")

	var x int = 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	const z int = 9
	// z = 10 // error
	fmt.Println(z)

}

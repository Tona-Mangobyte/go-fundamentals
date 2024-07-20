package main

import (
	"fmt"
	"math/rand"
)

// Blocks, Shadows, and Control Structures
func main() {
	fmt.Println("========Blocks, Shadows, and Control Structures========")
	// Blocks

	// Shadowing Variables
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	// Shadowing with multiple assignment
	x2 := 10
	if x2 > 5 {
		x2, y2 := 5, 20
		fmt.Println(x2, y2) // 5 20
	}
	fmt.Println(x2) // 10

	fmt.Println(true)
	true := 10
	fmt.Println(true)

	// If
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// Scoping a variable to an if statement
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// Out of scope
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	fmt.Println(n)

	// for
	fmt.Println("For loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	// The Condition-Only for Statement
	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}

	// The Infinite for Statement
	/*for {
		fmt.Println("Hello")
	}*/

	// break and continue
	/*do {
		// things to do in the loop
	} while (CONDITION);*/

	/*for {
		// things to do in the loop
		if !CONDITION {
			break
		}
	}*/

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// The for-range Statement
	fmt.Println("for-range")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Arrays
	var x [3]int
	x[0] = 5
	x[1] = 10
	x[2] = 15
	fmt.Println(x)

	var x2 = [3]int{10, 20, 30}
	fmt.Println(x2)

	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x3)

	var x4 = [...]int{10, 20, 30}
	fmt.Println(x4)

	var x5 = [...]int{1, 2, 3}
	var y5 = [3]int{1, 2, 3}
	fmt.Println(x5 == y5) // prints true

	// Slices
	x6 := []int{1, 2, 3, 4, 5}
	y6 := []int{1, 2, 3, 4, 5}
	z6 := []int{1, 2, 3, 4, 5, 6}
	// s6 := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x6, y6)) // prints true
	fmt.Println(slices.Equal(x6, z6)) // prints false
	// fmt.Println(slices.Equal(x6, s6)) // does not compile

	// len
	fmt.Println(len(z6))

	//append
	var x7 []int
	x7 = append(x7, 10) // assign result to the variable that's passed in
	fmt.Println(x7)

	var x8 = []int{1, 2, 3}
	x8 = append(x8, 4)
	fmt.Println(x8)

	y9 := []int{20, 30, 40}
	x8 = append(x8, y9...)
	fmt.Println(x8)

	//make
	x9 := make([]int, 5)
	fmt.Println(x9)
	x9 = append(x9, 10)
	fmt.Println(x9)

	x10 := make([]int, 5, 10)
	fmt.Println(x10)

	x11 := make([]int, 0, 10)
	x11 = append(x11, 5, 6, 7, 8)
	fmt.Println(x11)

	//Emptying a Slice
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))

	// copy
	x12 := []int{1, 2, 3, 4}
	y12 := make([]int, 4)
	num := copy(y12, x12)
	fmt.Println(y12, num)

	x13 := []int{1, 2, 3, 4}
	y13 := make([]int, 2)
	num2 := copy(y13, x13[2:])
	fmt.Println(y13, num2)

	// Maps
	var nilMap map[string]int
	fmt.Println(nilMap)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	ages := make(map[int][]string, 10)
	// ages[1] = []string{"Simple"}
	fmt.Println(ages)

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// Deleting from Maps
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	delete(m, "hello")
	fmt.Println(m)

	// Emptying a Map
	m2 := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m2, len(m2))
	clear(m2)
	fmt.Println(m2, len(m2))

	// Comparing Maps
	m3 := map[string]int{
		"hello": 5,
		"world": 10,
	}
	n3 := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(m3, n3)) // prints true

	// Using Maps as Sets
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	intSet2 := map[int]struct{}{}
	vals2 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals2 {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5 is in the set")
	}

	// Structs
	type person struct {
		name string
		age  int
		pet  string
	}
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)

	beth.name = "Bob" // A field in a struct is accessed with dot notation
	fmt.Println(beth.name)

	// Anonymous Structs
	var person2 struct {
		name string
		age  int
		pet  string
	}
	person2.name = "bob"
	person2.age = 50
	person2.pet = "dog"
	fmt.Println(person2)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)

	// Comparing and Converting Structs
	type firstPerson struct {
		name string
		age  int
	}
	type secondPerson struct {
		name string
		age  int
	}
	type thirdPerson struct {
		age  int
		name string
	}
	type fourthPerson struct {
		firstName string
		age       int
	}
	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}

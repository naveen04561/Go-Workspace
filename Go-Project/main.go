package main

import "fmt"

var d int

func main() {
	fmt.Println("Hello")
	var name string = "Naveen"
	fmt.Print(name + "\n")

	// Variables

	card := "Ace"
	card = "Diamond"
	fmt.Println(card)

	var p int
	var t = 2
	p = 0
	d = 40
	fmt.Println(p + t + d)

	newFunc()

	//fmt.Println(greet())

	// Slices and Loops

	nums := []int{1, 2}

	nums = append(nums, 4)

	for i, num := range nums {
		fmt.Println(i, num)
	}
}

// function

func newFunc() string {
	return "World"
}

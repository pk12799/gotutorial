package main

import "fmt"

func main() {
	//initilize array with direct intilizing
	first := [15]int{1, 2, 3, 45, 78}
	//intilize array through var keyword
	var second [15]int
	//copy arr array to brr array
	second = first
	// change the brr value
	second[4] = 10
	third := []int{10, 45}

	// println the array
	fmt.Println(second, first)
}

package main

import "fmt"

func main() {
	//working with slices
	//inilize slice with var keywork
	//slices are dynamic array
	var first = []int{1, 5, 45, 78, 59, 124, 369}
	//copy the first into second
	second := first
	//slice are working on pointer
	//if we change second slice element it will reflect in first also
	second[0] = 45
	/// we have append method for appending thge slice
	// here we have only append the second
	second = append(second, 10, 45, 45, 78)
	//we will copy the slice using copy method whose take two arguments first destination and second source
	third := []int{12, 785}
	copy(second, third)
	x := 5
	fmt.Println(second[1:x], first)
}

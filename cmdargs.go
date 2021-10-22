package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("gime some argument")
		os.Exit(1)
	}

	var max, min float64 = 0, 156897454
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		// fmt.Println(n)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

	}
	fmt.Println(min)
	fmt.Println(max)
}

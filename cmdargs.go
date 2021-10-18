// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) == 1 {
// 		fmt.Println("gime some argument")
// 		os.Exit(1)
// 	}
// 	// min := 0
// 	var max, min float64
// 	arguments := os.Args
// 	for i := 2; i < len(arguments); i++ {
// 		n, _ := strconv.ParseFloat(arguments[i], 64)
// 		if n < min {
// 			min = n
// 		}
// 		if n > max {
// 			max = n
// 		}

// 	}
// 	fmt.Println(min)
// 	fmt.Println(max)
// }
package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 100; x++ {
		// num := rand.Int()
		if x&1 == 1 {
			fmt.Printf("%d is odd\n", x)
		} else {
			fmt.Printf("%d is even\n", x)
		}
	}
}

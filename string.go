package main

import (
	"fmt"
	"strconv"
)

func alphasort(str []rune, depth int) {
	for x := range str {
		y := x + 1
		for y = range str {
			if str[x] < str[y] {
				str[x], str[y] = str[y], str[x]
			}
		}
	}
}

// func rtoc(r []rune) string {
// 	str := ""
// 	for _, c := range r {
// 		str += fmt.Sprintf("%c", c)
// 	}
// 	return str
// }

// func main() {
// 	str := "parvezkhan"
// 	runes := []rune(str)
// 	// fmt.Println(runes)
// 	alphasort(runes, 1)
// 	// rtoc(runes)
// 	fmt.Println(runes)
// 	fmt.Printf("%c", runes)
// 	fmt.Println()
// }

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0])
	fmt.Println(strconv.Itoa(11235489))
	var x string = "hgdfshdg"
	fmt.Println(x)
}

package main

import "fmt"

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
func main() {
	str := "parvez"
	runes := []rune(str)
	// fmt.Println(runes)
	alphasort(runes, 1)
	// rtoc(runes)
	fmt.Printf("%c", runes)
	fmt.Println()
}

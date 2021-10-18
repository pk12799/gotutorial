package main

import (
	"io"
	"os"
)

func main() {
	mystr := "hi"
	// at := os.Args[1:]
	// if len(at) == 1 {
	// 	mystr = "write something"
	// } else {
	// 	io.WriteString(os.Stdout, at[1])
	// }
	io.WriteString(os.Stdout, mystr)
	io.WriteString(os.Stdout, "\n")
}

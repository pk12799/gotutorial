package main

import (
	"fmt"
	"unsafe"
)

func main() {
	if 3 > 2 {
		fmt.Println("false")
		// fallthrough
		fmt.Println("hfhgddfg")
	} else {
		fmt.Println("true")
	}
	var a interface{} = 1
	switch a.(type) {
	case int:
		fmt.Println(unsafe.Sizeof(a))
		// fallthrough
	case float64:
		fmt.Println("float", unsafe.Sizeof(a))
		// fallthrough
	default:
		fmt.Println("none")
	}
}

// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func hello(i int) {
// // 	// time.Sleep(1 * time.Second)
// // 	fmt.Println("hello", i)

// // }
// // func hels(i int) {

// // 	fmt.Println("hels", i)

// // }
// // func hyts(i int) {
// // 	time.Sleep(1 * time.Second)
// // 	fmt.Println("hyts", i)
// // }
// // func main() {
// // 	for i := 0; i < 10; i++ {
// // 		go hello(i)
// // 		hyts(i)
// // 		go hels(i)
// // 		go hello(i)
// // 		//return

// // 	}
// // }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func numbers() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(12 * time.Millisecond)
// 		fmt.Printf("%d ", i)
// 	}
// }

// // func alphabets() {
// // 	for i := 'a'; i <= 'e'; i++ {
// // 		time.Sleep(400 * time.Millisecond)
// // 		fmt.Printf("%c ", i)
// // 	}
// // }
// func main() {
// 	go numbers()
// 	//go alphabets()
// 	time.Sleep(30 * time.Millisecond)
// 	fmt.Println("main terminated")
// }
package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello(i *int) {
	*i += 112
}
func hts(i *int) {
	*i += 10
	j := 10
	fmt.Println(j)
	t := 10
	switch t {
	case 10:
		fmt.Println("switch")
	case 20:
		fmt.Println("switch 20")
	default:
		fmt.Println("default")
	}
}
func f() {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print(i, " ")
	}
}
func main() {
	i, j := 10, 12
	p := &i
	q := &j

	go hello(p)
	f()
	go hts(q)

	// time.Sleep(1 * time.Second)
	fmt.Println("hello", *p, *q)
	runtime.GOMAXPROCS(120)
	fmt.Println("thereds", runtime.GOMAXPROCS(-1))
}

package main

import "fmt"

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func printChannelData(c chan int) {
// 	fmt.Println("Data in channel is: ", <-c)
// }
// func main() {
// 	fmt.Println("Main started...")
// 	//create channel of int
// 	d := make(chan int)
// 	c := make(chan int)
// 	// var c chan int

// 	// call to goroutine

// 	go printChannelData(c)
// 	go printChannelData(d)
// 	// put the data in channel

// 	c <- 10
// 	d <- 20
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main ended...")

// }
// func printh(m chan *int) {
// 	d := <-m
// 	fmt.Println("h", *d)
// }
// func main() {
// 	m := make(chan *int)
// 	s := 20
// 	go printh(m)
// 	m <- &s
// 	fmt.Println("hello")
// }
// func myfunc(ch chan int) {

// 	fmt.Println(234 + <-ch)
// }
// func main() {
// 	fmt.Println("start Main method")
// 	// Creating a channel
// 	ch := make(chan int)

// 	go myfunc(ch)
// 	ch <- 23
// 	fmt.Println("End Main method")
// }

// package main

// import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)

	// Anonymous goroutine
	go func() {

		mychnl <- "jhkj"
		mychnl <- "gfgds"
		close(mychnl)
		// mychnl <- "Gdjd"
		// mychnl <- "Ajjksj"

	}()

	// Using for loop
	// fmt.Println(<-mychnl)
	for res := range mychnl {
		fmt.Println(res)
	}
}

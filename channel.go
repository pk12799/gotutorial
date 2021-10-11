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
func printh(m chan *int) {
	d := <-m
	fmt.Println("h", *d)
}
func main() {
	m := make(chan *int)
	s := 20
	go printh(m)
	m <- &s
	fmt.Println("hello")
}

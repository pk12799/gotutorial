package main

import (
	"fmt"
	"sync"
)

func square(a chan int) {
	defer wg.Done()
	c := <-a
	c = c * c
	a <- c
}

func cube(b chan int) {
	c := <-b
	c = c * c * c
	b <- c
	// defer wg.Done()
}

func add(a chan int, b chan int, c chan int) {
	c <- <-a + <-b
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make(chan int)
	a := make(chan int)
	b := make(chan int)
	go square(a)
	a <- 2
	go cube(b)
	b <- 2
	go add(a, b, c)
	fmt.Println(<-c)
	wg.Wait()
}

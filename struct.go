package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type animal struct {
	res     func(int, int) int
	name    string
	origin  string
	country string
}
type birds struct {
	animal
	fly   bool
	speed float64
}
type cars struct {
	birds
	speed  float64
	model  string
	lounch string
}

func main() {
	a := animal{
		name:    "jhfh",
		origin:  "hgdf",
		country: "dsgfhf",
		res: func(ma int, tp int) int {
			a := ma * tp
			fmt.Println(a)
			return a
		},
	}
	b := birds{
		a,
		false,
		48.78,
	}
	fmt.Println(b, a)
	c := cars{

		birds:  b,
		model:  "fantom",
		lounch: strconv.Itoa(1999),
		speed:  1457,
	}
	d := animal{}
	fmt.Println(a.res(23, 34), d, unsafe.Sizeof(a), c)
}

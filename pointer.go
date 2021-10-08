package main

import (
	"fmt"
	"reflect"
)

type mstty struct {
	foo int
}

func zoo(w *int) {
	*w = 1 + 1 + 45878
}
func swap(m *int, n *int) {
	*m, *n = *n, *m
}

func main() {
	var ms *int
	my := *&mstty{}
	my.foo = 12
	fmt.Println(reflect.TypeOf(my.foo))
	var a **int
	a = &ms
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	x := new(int)
	c := x
	m := 12
	fmt.Println(reflect.TypeOf(x), x, reflect.TypeOf(c))
	zoo(&m)
	fmt.Println(*x)
	fmt.Println(reflect.TypeOf(m))
	m = 12
	n := 15
	swap(&m, &n)
	fmt.Println(m, n)
}

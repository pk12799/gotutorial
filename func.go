package main

import (
	"fmt"
	"reflect"
)

//return local variable as pointer
func zo(a *int, b int) (*int, error) {
	c := *a + b
	return &c, nil
}

//pass struct in func
type myst struct {
	name, country, origin string
}

//pass struct as reference and change the name value whose reflect on all scope
func (ms *myst) intro() {
	fmt.Println(*ms)
	ms.name = "pk"

}

func main() {
	a := 12
	b := 15
	d, err := zo(&a, b)
	if err != nil {
		return
	}
	fmt.Println(*d)
	ms := myst{
		name:    "parvez",
		country: "india",
		origin:  "asia",
	}
	//call rhe intro method with struct
	ms.intro()
	fmt.Println(ms)

	//anonymous function
	var di func(int, int) (int, error)
	di = func(a, b int) (int, error) {
		c := a + b
		return c, nil
	}
	f, er := di(12, 45)
	if er != nil {
		return
	}
	fmt.Println(f)
	//another way to intilize anonymous func
	func() {
		a := 78
		b := 459
		c := a + b

		fmt.Println(c)
	}()
	//assign a function into variable
	l := func(a ...float64) int {
		f := 45
		b := 78457
		c := f + b
		fmt.Println(a)
		return c
	}
	// v := []int{12, 45, 78, 98}
	z := l(12, 45, 78, 78.01)
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(l))
}

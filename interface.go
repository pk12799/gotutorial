package main

import "fmt"

type aabss interface {
	jgc(int, int) int
}
type str struct {
	a int
}

func (m str) jgc(a, b int) int {
	c := m.a + a + b
	return c
}
func main() {
	var d aabss
	b := &str{}
	b.a = 12
	c := d.jgc(12, b.a)
	fmt.Println(c, b.a)
}

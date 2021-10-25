package main

import "fmt"

type Maxheap struct {
	array []int
}

func (m *Maxheap) Insert(key int) {
	m.array = append(m.array, key)
	m.Maxheapifyup(len(m.array) - 1)
}
func (m *Maxheap) Extarct() int {
	extract := m.array[0]
	l := len(m.array) - 1
	if len(m.array) == 0 {
		fmt.Println("zero size")
		return -1
	}
	m.array[0] = m.array[l]
	m.array = m.array[:l]
	m.Maxheapifydown(0)

	return extract
}
func (m *Maxheap) Maxheapifydown(index int) {
	lastIndex := len(m.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l <= lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.array[index] < m.array[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}
func (m *Maxheap) Maxheapifyup(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
func (m *Maxheap) swap(i1, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}
func main() {
	max := &Maxheap{}
	buildheap := []int{10, 20, 56, 78, 12, 54, 487}
	for _, v := range buildheap {
		max.Insert(v)
		fmt.Println(max)
	}
	fmt.Println(max)
	for i := 0; i < 5; i++ {
		max.Extarct()
		fmt.Println(max)
	}
}

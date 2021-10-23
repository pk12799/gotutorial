package main

import "fmt"

const Arraysize = 7

type hashtab struct {
	array [Arraysize]*bucket
}
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//operations
func (h *hashtab) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}
func (h *hashtab) search(key string) {

}
func (h *hashtab) delete(key string) {

}
func Init() *hashtab {
	resu := &hashtab{}
	for i := range resu.array {
		resu.array[i] = &bucket{}
	}
	return resu
}

//bucket
func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Arraysize
}

func main() {
	Hash := Init()
	fmt.Println(Hash)
	k := hash("RANDY")
	fmt.Println(k)
	test := &bucket{}
	test.insert("RANDY")

}

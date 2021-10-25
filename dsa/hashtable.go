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
func (h *hashtab) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
func (h *hashtab) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
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
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "alredy exits")
	}
}
func (b *bucket) search(k string) bool {
	curNode := b.head
	for curNode != nil {
		if curNode.key == k {
			return true
		}
		curNode = curNode.next
	}
	return false

}

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
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
	// fmt.Println(Hash)
	list := []string{
		"24758",
		"ritik",
		"indra",
		"nikhil",
		"parvez Khan",
		"mantu",
		"pankaj",
		"satytam",
	}
	// test := &bucket{}
	for _, v := range list {
		Hash.Insert(v)
	}
	// fmt.Println(test.search("parvez"))
	// // test.insert("sANDY")
	// // test.delete("RANDY")
	fmt.Println(*Hash)
	fmt.Println(Hash.Search("parvez "))
}

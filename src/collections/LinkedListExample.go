package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	createList()

	kids := []Person{
		{"Jack", 23},
		{"Bill", 20},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	fmt.Println()

}

func createList() {
	var list list.List
	//var list list.New() //the same as above

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Person struct {
	Name string
	Age  int
}
type ByName []Person

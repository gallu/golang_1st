package main

import "fmt"

type KeyHoge struct {
	Id int
}
type KeyFoo struct {
	Name string
}

func (k KeyHoge) test() {
}

func (k KeyFoo) test() {
}

type KeyI interface {
	test()
}

// interface guard
var _ KeyI = (*KeyHoge)(nil)
var _ KeyI = (*KeyFoo)(nil)

func add(m map[KeyI]string, key KeyI, val string) map[KeyI]string {
	m[key] = val
	return m
}

func main() {
	m := make(map[KeyI]string, 0)
	add(m, KeyHoge{Id: 1}, "name")
	add(m, KeyFoo{Name: "key"}, "name 2")

	fmt.Printf("%#v \n", m)
}

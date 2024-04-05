package main

import "fmt"

type Hoge struct {
	List  map[string][]*Foo
	List2 map[string][]*Foo
}

type Foo struct {
	Data string
}

func NewHoge() *Hoge {
	return &Hoge{
		List:  map[string][]*Foo{},
		List2: map[string][]*Foo{},
	}
}

func NewFoo(s string) *Foo {
	return &Foo{
		Data: s,
	}
}

func listAdd(l map[string][]*Foo, key string, s string) {
	l[key] = append(l[key], NewFoo(s))
}

func main() {
	{
		h := NewHoge()
		key := "key 1"
		h.List[key] = append(h.List[key], NewFoo("d1"))
		h.List[key] = append(h.List[key], NewFoo("d2"))
		fmt.Printf("%+v \n", h)
	}
	{
		h := NewHoge()
		key := "key 1"
		listAdd(h.List, key, "d1")
		listAdd(h.List, key, "d2")
		listAdd(h.List2, key, "d111")
		listAdd(h.List2, key, "d222")

		fmt.Printf("%+v \n", h)
	}
	{
		h := NewHoge()
		key := "key 1"
		l := h.List
		listAdd(l, key, "d1")
		listAdd(l, key, "d2")
		l = h.List2
		listAdd(l, key, "d111")
		listAdd(l, key, "d222")
		fmt.Printf("%+v \n", h)
	}
}

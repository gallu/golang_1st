package main

import "fmt"

type LT map[string][]*Foo

type Hoge struct {
	List  LT
	List2 LT
}

type Foo struct {
	Name string
}

func NewFoo(n string) *Foo {
	return &Foo{
		Name: n,
	}
}

func Test(m LT, key string, value *Foo) {
	fmt.Printf("%p \n", m)
	_, ok := m[key]
	if !ok {
		m[key] = []*Foo{}
	}
	m[key] = append(m[key], value)
}

func main() {
	h := &Hoge{
		List:  LT{},
		List2: LT{},
	}
	fmt.Printf("%p \n", h.List)
	Test(h.List, "t1", NewFoo("t1t1"))
	Test(h.List, "t2", NewFoo("t2t2"))
	fmt.Printf("%p \n", h.List)

	fmt.Println("")

	fmt.Printf("%p \n", h.List2)
	Test(h.List2, "t1", NewFoo("t1t1"))
	Test(h.List2, "t2", NewFoo("t2t2"))
	fmt.Printf("%p \n", h.List2)

}

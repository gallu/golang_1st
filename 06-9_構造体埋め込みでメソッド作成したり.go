package main

import "fmt"

type Hoge struct {
	Data []int
}
type Foo struct {
	Id int
	Hoge
}

func NewFoo(id int) *Foo {
	return &Foo{
		Id: id,
		Hoge: Hoge{
			Data: make([]int, 0),
		},
	}
}

func (h *Hoge) Push(i int) {
	h.Data = append(h.Data, i)
}

func (f *Foo) GetHoge() *Hoge {
	return &f.Hoge
}

func main() {
	// これはいける
	{
		f := NewFoo(1)
		f.Hoge.Push(1)
		f.Hoge.Push(2)
		f.Hoge.Push(3)
		fmt.Printf("%+v \n", f)
	}
	// あぁやっぱりこれもいけるのか
	{
		f := NewFoo(1)
		f.Push(11)
		f.Push(22)
		f.Push(33)
		fmt.Printf("%+v \n", f)
	}
	// ポインタげとってもいけるのね
	{
		f := NewFoo(1)
		h := f.GetHoge()
		h.Push(111)
		h.Push(222)
		h.Push(333)
		fmt.Printf("%+v \n", f)
	}
}

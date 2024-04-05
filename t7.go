package main

import "fmt"

type Hoge struct{}

func (h Hoge) Func1() {
	fmt.Println("Func1")
}
func (h Hoge) Func2() {
	fmt.Println("Func2")
}

type HogeI interface {
	Func1()
}

func NewHoge() HogeI {
	return &Hoge{}
}

func main() {
	h := NewHoge()
	h.Func1()
	// h.Func2()

	// 型アサーションで型を変更すると使える
	hh, ok := h.(*Hoge)
	if !ok {
		panic("だめぽ")
	}
	hh.Func2()
}

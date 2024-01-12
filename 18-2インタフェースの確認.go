package main

import "fmt"

type Hoge struct{}

func (h *Hoge) func1() {
	fmt.Println("func1")
}
func (h *Hoge) func2() {
	fmt.Println("func2")
}
func (h *Hoge) func3() {
	fmt.Println("func3")
}

type HogeI1 interface {
	func1()
}
type HogeI2 interface {
	func2()
}

func NewHoge() *Hoge {
	return &Hoge{}
}
func NewHoge1() HogeI1 {
	return &Hoge{}
}
func NewHoge2() HogeI2 {
	return &Hoge{}
}

func main() {
	h := NewHoge()
	h.func1()
	h.func2()
	h.func3()

	h1 := NewHoge1()
	h1.func1()
	// h1.func2() // だめ

	h2 := NewHoge2()
	// h2.func1() // だめ
	h2.func2()
}

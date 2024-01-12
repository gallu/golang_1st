package main

import "fmt"

/* 以下、なんか違うパッケージにある想定 */
type Hoge struct {
}

type LogicI interface {
	FuncA()
}

// ガード
var _ LogicI = (*Hoge)(nil)

func NewHoge() LogicI {
	return &Hoge{}
}

func (h *Hoge) FuncA() {
	fmt.Println("FuncA")
}
func (h *Hoge) FuncB() {
	fmt.Println("FuncB")
}

/* 以上、なんか違うパッケージにある想定 */

func main() {
	{
		h := NewHoge()
		fmt.Printf("%T \n", h) // *main.Hoge
		h.FuncA()
		// h.FuncB() // これはcallできない
	}
	{
		h := &Hoge{}
		fmt.Printf("%T \n", h) // *main.Hoge
		h.FuncA()
		h.FuncB()
	}
}

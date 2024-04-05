package main

import "fmt"

type Hoge struct {
}

func NewHoge() *Hoge {
	return &Hoge{}
}

func (h *Hoge) Func1(i int) {
	fmt.Println("Func1", i)
}
func (h *Hoge) Func2(i int) {
	fmt.Println("Func2", i)
}
func (h *Hoge) Func3(i int) {
	fmt.Println("Func3", i)
}

func main() {
	//
	{
		h := NewHoge()
		h.Func1(1)
		h.Func2(2)
		h.Func3(3)
	}
	//
	{
		h := NewHoge()
		f := h.Func1
		f(1)
	}
	//
	{
		h := NewHoge()
		arr := []func(int){
			h.Func1,
			h.Func2,
			h.Func3,
		}
		for i, f := range arr {
			f(i)
		}
	}

}

package main

import "fmt"

type Detail struct {
	Num int
}
type Body1 struct {
	Detail
}
type Body2 struct {
	Detail
}

func (h *Body1) fn() {
}
func (f *Body2) fn() {
}

type BodyI interface {
	fn()
}

func (d *Detail) EffNum(i int) {
	d.Num += i
}

func main() {
	//
	{
		b1 := &Body1{
			Detail{
				Num: 0,
			},
		}
		fmt.Printf("%+v \n", b1)

		b1.EffNum(123)
		fmt.Printf("%+v \n", b1)
	}
	{
		b1 := &Body2{
			Detail{
				Num: 0,
			},
		}
		fmt.Printf("%+v \n", b1)

		b1.EffNum(123)
		fmt.Printf("%+v \n", b1)
	}

}

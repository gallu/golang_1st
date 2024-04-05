package main

import "fmt"

type Hoge struct {
	Id int
}

func NewHoge(i int) **Hoge {
	r := &Hoge{Id: i}
	return &r
}

type HogeList []PPHoge

// type HogeList []**Hoge

type PPHoge **Hoge

func (h *Hoge) t() {
	h.Id *= 2
	fmt.Println("func", h.Id)
}

func loop(l HogeList) {
	for k, v := range l {
		(*v).t()
		fmt.Printf("%d, %+v \n", k, *v)
		fmt.Printf("%d, %+v \n", k, **v)
	}
}

func main() {
	{
		// l := HogeList{}
		// l := []**Hoge{} // エラーになると思ったけど、ならんねぇ
		l := []PPHoge{} // エラーになると思ったけど、ならんねぇ
		l = append(l, NewHoge(1))
		l = append(l, NewHoge(2))
		l = append(l, NewHoge(3))
		loop(l)
	}
}

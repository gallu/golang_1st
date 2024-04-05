package main

import "fmt"

type Hoge struct {
	id   int
	next *Hoge
}

func NewHoge(i int, n *Hoge) *Hoge {
	return &Hoge{
		id:   i,
		next: n,
	}
}

func Use(h *Hoge) *Hoge {
	fmt.Printf("%#v \n", h)
	return h.next
}

func main() {
	{
		/*
			h3 := NewHoge(3, nil)
			h2 := NewHoge(2, h3)
			h1 := NewHoge(1, h2)
		*/
		//
		h1 := NewHoge(1, nil)
		//
		h2 := NewHoge(2, nil)
		h1.next = h2
		//
		h3 := NewHoge(3, nil)
		h2.next = h3

		H := h1
		for H != nil {
			H = Use(H)
		}
	}
}

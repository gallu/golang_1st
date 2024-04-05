package main

import "fmt"

type Hoge struct {
	id int
	h  []Hoge
}

func main() {
	{
		h := Hoge{
			id: 10,
			h:  nil,
		}
		fmt.Printf("%#v \n", h)
	}
	{
		h := Hoge{
			id: 10,
			h: []Hoge{
				Hoge{
					id: 2,
					h:  nil,
				},
				Hoge{
					id: 3,
					h:  nil,
				},
			},
		}
		fmt.Printf("%#v \n", h)
	}
}

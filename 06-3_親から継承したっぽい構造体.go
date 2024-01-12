package main

import (
	"encoding/json"
	"fmt"
)

type IdInt int
type IdStr string

type Hoge struct {
	Id   int
	Name string
}

type Foo struct {
	Id   int
	Type string
	Num  float64
}

type Bar struct {
	Id   int
	Days int
}

// カテゴライズするためのダミーメソッド
func (h *Hoge) Ttt() {}
func (h *Foo) Ttt()  {}
func (h *Bar) Ttt()  {}

type DetailI interface {
	Ttt()
}

// interface guard
var _ DetailI = (*Hoge)(nil)
var _ DetailI = (*Foo)(nil)
var _ DetailI = (*Bar)(nil)

type Total struct {
	Detail []DetailI
}

func NewTotal() *Total {
	return &Total{}
}
func (t *Total) AddDetail(d DetailI) {
	if t.Detail == nil {
		t.Detail = make([]DetailI, 0)
	}
	t.Detail = append(t.Detail, d)
}

func main() {
	//
	{
		t := NewTotal()
		t.AddDetail(&Hoge{
			Id:   1,
			Name: "Hoge name",
		})
		t.AddDetail(&Foo{
			Id:   2,
			Type: "type",
			Num:  3.14,
		})
		t.AddDetail(&Bar{
			Id:   3,
			Days: 999,
		})

		s, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s))
	}
}

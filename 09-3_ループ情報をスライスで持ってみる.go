package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hoge struct {
	data []string
	pos  []string
}

func NewHoge(data []string) *Hoge {
	return &Hoge{
		data: data,
	}
}

// 開始の合図
func (h *Hoge) Start() {
	h.pos = make([]string, len(h.data))
	copy(h.pos, h.data)
}
func (h *Hoge) Next() bool {
	return len(h.pos) != 0
}
func (h *Hoge) Get() string {
	// データがなければパニックでも
	if len(h.pos) == 0 {
		panic("ほげら")
	}
	// XXX ここでsortしてもよい
	sort.SliceStable(h.pos, func(i, j int) bool {
		// return h.pos[i] > h.pos[j]
		i = rand.Intn(999999)
		j = rand.Intn(999999)
		return i > j
	})

	// 取り出して
	r := h.pos[0]
	// 先頭の要素を削って
	h.pos = h.pos[:0+copy(h.pos[0:], h.pos[0+1:])]

	// 削った値をreturnする
	return r
}

func main() {
	{
		h := NewHoge([]string{
			"hoge",
			"foo",
			"bar",
		})

		for i := 0; i < 5; i++ {
			h.Start()
			for h.Next() {
				s := h.Get()
				fmt.Println(s)
			}
			fmt.Println("")
		}
	}
}

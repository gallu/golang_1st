package main

import "fmt"

type HogeStruct struct {
	id int
}

// これはいわゆる「DBで複数レコード取得」とかのパターン(今回の範疇外なので軽く書いておくだけ)
type HogeSlice []*HogeStruct

type keyId int
type HogeMap map[keyId][]*HogeStruct

func main() {
	//
	{
		h := make(HogeMap, 0)
		add(h, 1, &HogeStruct{id: 1})
		add(h, 2, &HogeStruct{id: 2})
		add(h, 1, &HogeStruct{id: 3})

		fmt.Printf("%#v\n", h)
	}
}

// XXX mapはそもそもポインタっぽいので破壊的変更になる感じ？？？(この辺まだ把握しきれてない)
func add(h HogeMap, k keyId, hs *HogeStruct) {
	_, ok := h[k]
	if !ok {
		h[k] = make([]*HogeStruct, 0, 0)
	}
	h[k] = append(h[k], hs)
}

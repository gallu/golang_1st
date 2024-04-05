package main

import (
	"container/list"
	"fmt"
)

type Hoge struct {
	Id int
}

func NewHoge(i int) *Hoge {
	return &Hoge{
		Id: i,
	}
}

func main() {
	{
		l := list.New()

		e1 := l.PushBack(NewHoge(1))
		fmt.Printf("e1 %#v\n", e1.Value)

		l.PushBack(NewHoge(2))

		e3 := l.PushFront(NewHoge(3))
		fmt.Printf("e3 %#v\n", e3.Value)

		l.PushFront(NewHoge(4))

		l.InsertBefore(NewHoge(5), e1)
		l.InsertAfter(NewHoge(6), e1)

		// 一覧表示
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%T, %#v\n", e.Value, e.Value)
		}

		// findして削除
		for e := l.Front(); e != nil; e = e.Next() {
			// h := e.Value.(*Hoge)
			// fmt.Println(h.Id)
			if e.Value.(*Hoge).Id == 1 {
				l.Remove(e)
			}
		}

		// 改めて一覧表示
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("\t%#v\n", e.Value)
		}
	}

	// XXX ジェネリクスで型指定したlistはまだないぽ(Go 2で出てくる？)
}

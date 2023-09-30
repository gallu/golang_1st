package main

import "fmt"

type Stack[T any] struct {
	v []T
}

func New[T any]() Stack[T] {
	r := Stack[T]{}
	return r
}

func (s *Stack[T]) Push(x T) {
	s.v = append(s.v, x)
}

func (s *Stack[T]) Pop() (T, bool) {
	// 長さが0ならNGに
	if len(s.v) == 0 {
		// もうちょっと楽に書ける方法があったと思うんだけど……
		var zero T
		return zero, false
	}
	// 末尾要素、これで取るしかないのか?
	r := (*s).v[len((*s).v)-1]
	s.v = s.v[:len(s.v)-1]
	return r, true
}

func main() {
	{
		arr := New[int]()
		fmt.Println(arr)

		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		fmt.Println(arr)

		for {
			r, ok := arr.Pop()
			if ok == false {
				fmt.Println("おわり!！")
				break
			}
			fmt.Println(r)
			fmt.Println(arr)
		}
	}
}

package main

import "fmt"

// 構造体にもジェネリクス使えるぽ
type SHoge[T any] struct {
	Id  int
	Msg *T
}

// 「構造体のジェネリクス」使うとき、たぶん割と鉄板な書き方なんだろうなぁ
func (sh SHoge[_]) getMsg() any {
	return sh.Msg
}

func NewSHoge[T any](id int, msg *T) *SHoge[T] {
	return &SHoge[T]{
		Id:  id,
		Msg: msg,
	}
}

func main() {
	msg := "message string"
	sh := NewSHoge(1, &msg)

	m := sh.getMsg()
	fmt.Println(*(m.(*string)))         // 個人的にはこっちが好きなんだけど
	fmt.Println(*m.(*string))           // golangらしい、だと多分こっちなんだろうなぁ
	fmt.Println(*sh.getMsg().(*string)) // 「中間変数も使わず」で、こっちかな？？？
}

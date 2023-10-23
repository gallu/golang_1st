package main

import "fmt"

// 埋め込む構造体
type baseAccount struct {
	id int
}

func (ba baseAccount) getId() int {
	return ba.id
}

// 埋め込み先構造体(フィールド名なし)
// XXX 透過的に使えるから、多分、こっちのほうがよさげ
type account struct {
	baseAccount
	name string
}

// 埋め込み先構造体(フィールド名あり)
type account2 struct {
	ba   baseAccount
	name string
}

func main() {
	// フィールド名なしの実験
	{
		ac := new(account)
		ac.id = 1
		ac.baseAccount.id = 123 // こっちでもいいぽい
		ac.name = "gallu"
		fmt.Println(ac.getId())
	}
	// フィールド名ありの実験
	{
		ac := new(account2)
		// ac.id = 1 // これはダメ(解決できないぽい)
		ac.ba.id = 1
		ac.name = "gallu"
		// fmt.Println(ac.getId()) // これもダメ
		fmt.Println(ac.ba.getId())
	}
}

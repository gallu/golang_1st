package main

import "fmt"

// 定数の宣言
// 先頭小文字なんで privateって感じ
const s string = "const string"

// 複数
// 型なし
// 先頭大文字なんでpublicって感じ
const (
	Id   int = 123
	Name     = "hogera"
)

func main() {
	fmt.Println(s)
	fmt.Println(Id, Name)
}

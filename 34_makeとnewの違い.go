package main

import "fmt"

func main() {
	// makeは「スライス」「マップ」「チャンネル」くらいでしか使わない(使えない?)
	arr := make([]int, 10)
	// newは「ポインタ型」を生成する。実質的には構造体に使うくらい?
	pi := new(int)
	*pi = 10

	fmt.Println(arr)
	fmt.Println(pi)  // こっちだとポインタ
	fmt.Println(*pi) // こっちだと実体
}

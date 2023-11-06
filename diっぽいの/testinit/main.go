package main

import (
	"fmt"
	"testinit/usecase"
)

func main() {
	// 通常の呼び出し
	{
		fmt.Println("通常の呼び出し")
		usecase.Hoge(10)
	}
	// (主にテスト想定)DIで依存性ぶちこんでの呼び出し
	{
		fmt.Println("テスト用の呼び出し")
		usecase.HogeTest(10)
	}
}

package main

import "fmt"

type iHoge int

func main() {
	//
	{
		var i int = int(-10)
		var i64 int64 = int64(i)
		var i8 int8 = int8(i)
		fmt.Println(i, i64, i8)
	}
	//
	{
		var i int = int(-10)
		var i64 uint64 = uint64(i) // オーバフロー(アンダーフロー)する
		var i8 uint8 = uint8(i)    // オーバフロー(アンダーフロー)する
		fmt.Println(i, i64, i8)
	}
	//
	{
		var i uint = uint(10)
		var i64 int64 = int64(i)
		var i8 int8 = int8(i)
		fmt.Println(i, i64, i8)
	}
	//
	{
		var i iHoge = -10
		fmt.Println(i)
		// i += 10 // これは出来ない
		// こっちはいける
		var ii int = int(i) + 30
		fmt.Println(ii)
		// これもいける
		i = iHoge(int(i) + 44)
		fmt.Println(i)
	}
}

package main

import "fmt"

func main() {
	// xを定義
	x := 1
	fmt.Println(x)

	// ブロック
	{
		// ブロックの中で普通にxを使う
		x = 2
		fmt.Println("\t", x)
	}
	// 普通にブロックの中の変更は受け継がれる
	fmt.Println(x)

	// ブロック
	{
		// ブロックの中で「新しく」変数をつくる
		// と、元のxが隠れる。これがシャドーイング
		var x int = 999
		fmt.Println("\t", x)
		fmt.Printf("\t%p \n", &x)
	}
	fmt.Println(x)
	fmt.Printf("%p \n", &x)

	{
		// 識別子にも上書き出来る事があるから注意!!
		fmt.Println(true)
		true := 99999
		fmt.Println(true)
	}
	// ブロックの外に出たら戻るけどね
	fmt.Println(true)
}

package main

import (
	"fmt"
)

type SHoge struct {
	Id   int
	Name string
}

func main() {
	// anyの時に「この型、って前提でいれる～」の指定が出来るぽ
	{
		var i any = 123
		j := i.(int)
		fmt.Println(j)
	}
	//
	{
		/*
			var i any = 123
			j := i.(int64) // 型が違うからpanicが発生する
			fmt.Println(j)
		*/
	}
	//
	{
		var i any = 123
		j, ok := i.(int64) // 型が違う、けどカンマOKイディオムでひっかけられる
		if !ok {
			fmt.Println("アサーションNG!!")
		} else {
			fmt.Println(j)
		}
	}
	// アサーション、ようは「モワっとさせた型を元に戻す」的な？
	{
		var sha any = SHoge{
			Id:   1,
			Name: "hoge",
		}
		fmt.Println(sha) // 一応出力はできるぽ
		// sha.Id = 999 // 「構造体としての操作」はでけない
		sh, ok := sha.(SHoge)
		if !ok {
			fmt.Println("アサーションNG!!")
		} else {
			sh.Id = 999 // こっちは当然、できる
			fmt.Println(sh)
		}
	}

	// 型switch
	{
		var i any = "str"
		// fmt.Println(i.(type)) " use of .(type) outside type switch" だって!!
		switch i.(type) {
		case nil:
			fmt.Println("nilだよ")
		case string:
			fmt.Println("文字だよ")
		case int, float64, float32:
			fmt.Println("数値だよ")
		default:
			fmt.Println("知らんがな")
		}
	}
}

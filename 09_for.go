package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 「普通」のfor文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// whileっぽい書き方
	{
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	}

	// 無限ループはコメントだけ
	/*
		for {
			// 処理
		}
	*/

	// 文字列とかスライスとかマップとかのranged for
	{
		// スライスのfor
		arr := []int{11, 22, 33, 44}
		for i, v := range arr {
			fmt.Println(i, v)
		}

		// mapのfor
		marr := map[string]string{
			"key":  "value",
			"key2": "value2",
		}
		for k, v := range marr {
			fmt.Println(k, v)
		}

		// 文字のfor
		s := "abcあいう"
		for i, v := range s {
			fmt.Println(i, v, string(v))
		}
		// あぁまぁいけるか
		for i, v := range "いろは" {
			fmt.Println(i, v, string(v))
		}
	}

	// rangeの値はcopyであって参照じゃないよ
	{
		// スライスのfor
		arr := []int{11, 22, 33}
		fmt.Println(arr)
		for i, v := range arr {
			v += v
			fmt.Println(i, v)
		}
		fmt.Println(arr)
	}

	// ラベル使っての多段ぶち抜きのcontinueとbreak
	{
	loop_label:
		for {
			fmt.Print("*")
			for {
				fmt.Print("+")
				if n := rand.Intn(5); n == 1 {
					fmt.Println("")
					break // こっちは1段
				}
				if n := rand.Intn(5); n == 0 {
					break loop_label // こっちはぶち抜き
				}
			}
		}
		fmt.Println("")
	}
	// ラベルの多段。予想通りの動き
	{
	loop_label_1:
		for {
			fmt.Print("*")
		loop_label_2:
			for {
				fmt.Print(":")
				for {
					fmt.Print("+")
					if n := rand.Intn(3); n == 0 {
						fmt.Print("?", n)
						break loop_label_2
					}
					if n := rand.Intn(5); n == 4 {
						fmt.Print(n)
						break loop_label_1
					}
				}
			}
		}
		fmt.Println("")
	}

}

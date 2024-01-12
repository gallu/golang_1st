package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	{
		// 条件式付近に文が書けるのは興味深い
		// 流石に2文(以上)は書けない 	// if a := 10; n := 20; n === 0
		if n := 1; n == 1 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
		// ブロック外だから怒られる
		// fmt.Println(n)

		rand.Seed(time.Now().UnixMicro()) // 1.20から非推奨らしいんだけど、動かしてる環境が1.19なので、とりあえず入れておく。これより crypto.rand のほうがいい気はする、けど「重い」らしいので今度確認
		if n := rand.Intn(5); n == 0 {
			fmt.Println(0)
		} else if n == 1 { // else if は離すタイプ
			fmt.Println(1)
		} else {
			fmt.Println(false)
		}

	}

	{
		b := true
		// これ、どっちでもよいぽ？
		if b == true {
			fmt.Println("ok")
		}
		if b {
			fmt.Println("ok")
		}
	}

	{
		/*
			// non-boolean condition in if statement がちゃんと出る！ 暗黙の変換とかない！！
			i := 1
			if i {
				fmt.Println("を？")
			}
		*/
	}
}

package main

import (
	"errors"
	"fmt"
)

type TestError1 struct {
	message string
}

func (te TestError1) Error() string {
	return "てすと1のえらら: " + te.message
}

var IsUseTestError1 TestError1 = TestError1{message: "よびよび"}

func errorTest1() (int, error) {
	// return 0, TestError1{message: "よびよび"} // こっちだとIs成功
	return 0, TestError1{message: "よびよびび?"} // こっちだとIs失敗
}

func main() {
	// 前座的に軽く
	{
		_, err := errorTest1()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Isは「エラーの(中の)値を比較」する
	{
		_, err := errorTest1()
		if err != nil {
			if errors.Is(err, IsUseTestError1) {
				fmt.Println("Is 成功!!") // こっち
			} else {
				fmt.Println("Is 失敗 orz")
			}
		}
	}
	{
		// 一応確認してみたい
		err := errors.New("てすと1のえらら: " + "よびよび")
		// fmt.Println(err)
		// fmt.Println(IsUseTestError1)
		// 流石にこれは駄目ぽい。ので多分「型と値の一致」なんだと思われる
		if errors.Is(err, IsUseTestError1) {
			fmt.Println("Is その2 成功!!")
		} else {
			fmt.Println("Is その2 失敗 だ～よ～ね～") // こっち
		}
	}

	// Asは「エラーの型を比較」する
	{
		var te1 TestError1
		_, err := errorTest1()
		if errors.As(err, &te1) {
			//fmt.Println(err)
			fmt.Println("As 成功!!") // こっち
		} else {
			fmt.Println("As 失敗 orz")
		}
	}
	// 「直接型名」でもいけるぽい
	// XXX 好みはこっちだなぁ
	{
		_, err := errorTest1()
		if errors.As(err, &TestError1{}) {
			//fmt.Println(err)
			fmt.Println("As その2 成功!!") // こっち
		} else {
			fmt.Println("As その2 失敗 orz")
		}
	}
}

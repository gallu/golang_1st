package main

import "fmt"

type BitParams int

const (
	Hoge BitParams = 1 << iota
	Foo
	Bar
	Piyo
)

func Check(params BitParams) (result string) {
	// 各パラメタの判定
	if params&Hoge == Hoge {
		result += "Hoge"
	}
	if params&Foo == Foo {
		result += "Foo"
	}
	if params&Bar == Bar {
		result += "Bar"
	}
	if params&Piyo == Piyo {
		result += "Piyo"
	}

	return result
}

func main() {
	// Hogeを判定
	{
		s := Check(Hoge)
		fmt.Println(s)
		if s != "Hoge" {
			fmt.Println("Hogeの判定が正しくありません")
		}
	}
	// HogeとFooを判定
	{
		s := Check(Hoge | Foo)
		fmt.Println(s)
		if s != "HogeFoo" {
			fmt.Println("HogeとFooの判定が正しくありません")
		}
	}
	// HogeとBarを判定
	{
		s := Check(Hoge | Bar)
		fmt.Println(s)
		if s != "HogeBar" {
			fmt.Println("HogeとBarの判定が正しくありません")
		}
	}
}

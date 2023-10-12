package main

import "fmt"

func hoge(i int, j int) string {
	s := fmt.Sprintf("%d + %d", i, j)
	return s
}
func foo(i int, j int) string {
	s := fmt.Sprintf("%d - %d", i, j)
	return s
}

// 「関数ポインタをreturnする」やつ
func getFunc() func(int) string {
	return func(i int) string {
		return fmt.Sprintf("param is %d", i)
	}
}

// 型宣言するともうちょっと綺麗。C の typedef だよねぇ
type opFuncP func(int, int) string

func main() {
	{
		// 関数ポインタ的なの
		var fn func(int, int) string = hoge // := でいいんだけど「あえて」綺麗に書いてみる。"func(int, int) string" までが型情報っぽ
		s := fn(10, 20)
		fmt.Println(s)

		fn = foo
		s = fn(10, 20)
		fmt.Println(s)
	}

	// typeで楽にしたのを使う
	{
		// 関数ポインタ的なの
		var fn opFuncP = hoge
		s := fn(10, 20)
		fmt.Println(s)
	}

	// 無名関数
	{
		func(i int) {
			fmt.Println(i * 2)
		}(11)

		// 変数に代入
		var fn func(int) = func(i int) {
			fmt.Println(i * 99)
		}
		fn(123)
	}

	// 「戻り値の関数ポインタ」を使う
	{
		fn := getFunc()
		s := fn(123)
		fmt.Println(s)
	}

	// fibonacciで確認。クロージャの時、変数は「同じ実体」を見ているっぽい。
	{
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
	}
}

// 必要ない宣言な気もするけど、まぁ勉強かねて
type fTest func() int

func fibonacci() fTest {
	a, b := 1, 0
	return func() int {
		fmt.Printf("%p \n", &a)
		a, b = b, a+b
		return a
	}
}

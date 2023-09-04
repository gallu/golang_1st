package main

import "fmt"

func main() {
	// 数値はアンダースコアで目印付けられる
	var i int = 1_234
	fmt.Println("i is ", i)

	// 2進数、8進数、16進数
	i = 0b00001111
	fmt.Println("i(2) is ", i)
	i = 0o777
	fmt.Println("i(8) is ", i)
	i = 0xff
	fmt.Println("i(16) is ", i)

	// 小数点数は e表記できる
	var f float64 = 123e2
	fmt.Println("f is ", f)
	f = 123e-2
	fmt.Println("f is ", f)

	// runeリテラルは「1文字」って意味合いっぽ
	var c rune = 'あ'
	fmt.Println("c is ", c, string(c)) // string()通さないと数値になっちゃうお

	// ダブルクォートは「改行コード入れられない」から注意
	s := "ab \n cd"
	fmt.Println("s is ", s)
	// バックスラッシュだと改行込み込みに出来る
	s = `abc
	def
	eee`
	fmt.Println("s is ", s)

	// bool
	var b bool = true
	fmt.Println("b is ", b)

	// 型変換 は 型名(変数) でやるぽ
	f = 123
	i = int(f)
	fmt.Println("i is ", i)

	// 連続した変数の宣言
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)
	{
		// ブロックの中は別スコープ
		x, y := 111, 222     // これによって上のx, yが隠れる(シャドーイングって言うぽい)
		fmt.Println(x, y, z) // zだけお外の
	}
}

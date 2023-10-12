package main

import (
	"fmt"
	"strconv"
)

func main() {
	{
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

		// 「数値を文字にする」時はこんな感じで
		{
			i := 123
			si := strconv.FormatInt(int64(i), 10) // intでは駄目ぽい
			fmt.Printf("%T, %s \n", si, si)

			// 小数点、ちょいと面倒ぽ(printfのほうがわかりやすそうな???)
			f = 123456789.1415926535
			sf := strconv.FormatFloat(f, 'f', -1, 64) // 精度とか無指定
			fmt.Printf("%T, %s \n", sf, sf)
			sf = strconv.FormatFloat(f, 'e', -1, 64) // 指数表記指定
			fmt.Printf("%T, %s \n", sf, sf)
			sf = strconv.FormatFloat(f, 'f', 2, 64) // 小数点第二位まで
			fmt.Printf("%T, %s \n", sf, sf)
			sf = strconv.FormatFloat(f, 'e', 2, 64) // 小数点第二位まで、指数表記
			fmt.Printf("%T, %s \n", sf, sf)
			sf = strconv.FormatFloat(f, 'g', -1, 64) // ある程度以上おおきな数値なら指数表記
			fmt.Printf("%T, %s \n", sf, sf)
			sf = strconv.FormatFloat(f, 'g', 2, 64) // 小数点第二位まで、ある程度以上おおきな数値なら指数表記
			fmt.Printf("%T, %s \n", sf, sf)
		}

		// 「文字を数値にする」時はこんな感じで
		{
			// ざっくり
			si := "123"
			i, err := strconv.Atoi(si)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%T, %d \n", i, i)

			// 基数とかbit長とか色々指定してきめ細かく
			i2, err := strconv.ParseInt(si, 10, 64)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%T, %d \n", i2, i2)

			// 小数点数は細かくしかないっぽ
			sf := "3.141592"
			f, err := strconv.ParseFloat(sf, 64)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%T, %f \n", f, f)
		}

		// 連続した変数の宣言
		var x, y, z int = 1, 2, 3
		fmt.Println(x, y, z)
		{
			// ブロックの中は別スコープ
			x, y := 111, 222     // これによって上のx, yが隠れる(シャドーイングって言うぽい)
			fmt.Println(x, y, z) // zだけお外の
		}
	}

	// zero値の確認
	{
		var i int
		fmt.Println("zero int", i)

		var f float64
		fmt.Println("zero float64", f)

		var s string
		fmt.Println("zero string", s)

		var b bool
		fmt.Println("zero bool", b)
	}
}

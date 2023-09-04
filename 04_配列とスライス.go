package main

import (
	"fmt"
)

func main() {
	// 配列は固定長
	{
		var arr [3]int = [3]int{1, 2, 3}
		fmt.Println(arr)
	}
	// 数とか適当に
	{
		var arr = [...]int{1, 2, 3}
		fmt.Println(arr)
	}

	// スライスだと可変長にできる
	{
		var arr = []int{1, 2, 3}
		fmt.Println(arr, len(arr), cap(arr))

		// 後ろに追加
		arr = append(arr, 444, 555)
		fmt.Println(arr, len(arr), cap(arr))
		// 前に追加
		// スライスの後ろに三点リーダ付けて中味展開する
		arr = append([]int{77, 88}, arr...)
		fmt.Println(arr, len(arr), cap(arr))
	}
	// make使う(キャパシティ指定できるから多分こっちが好み)
	{
		// 空っぽだけどキャパ20
		arr := make([]int, 0, 20)
		fmt.Println(arr, len(arr), cap(arr))
	}
	// 長さだけ指定すると「ゼロ値」が入る
	{
		arr := make([]int, 10, 10)
		fmt.Println(arr, len(arr), cap(arr))
	}
	// スライスの部分的な抜き出し
	// XXX イメージとしては、C++のSTLのbegin()とend()の関係っぽい(end()になったら終了、ってあたりが)
	{
		arr := []int{0, 1, 2, 3, 4, 5}
		x := arr[:]
		fmt.Println(":", x)
		x = arr[1:2]
		fmt.Println("1:2", x)
		x = arr[1:]
		fmt.Println("1:", x)
		x = arr[:3]
		fmt.Println(":3", x)
		// スライスは「同じ領域を見ている」っぽい
		x[1] = 11111
		fmt.Println(arr)
		fmt.Printf("%p \n", arr)
		fmt.Printf("%p \n", &arr[0])
		fmt.Printf("%p \n", x)
		fmt.Printf("%p \n", &x[0])
		// 「途中抜き出し」だとその分ポインタ進んでる(けどメモリ領域はかぶってるよねぇ)
		x = arr[1:2]
		fmt.Println("1:2", x)
		fmt.Printf("%p \n", arr)
		fmt.Printf("%p \n", &arr[1])
		fmt.Printf("%p \n", x)

		// フルスライス(キャパシティを指定するスライス)
		fmt.Println(arr, len(arr), cap(arr))
		// x = arr[0:5:7] // これがフルスライス。ただ「元のスライスのキャパ」を超える値は指定できない
		x = arr[0:5:cap(arr)] // これがフルスライス。ただ「元のスライスのキャパ」を超える値は指定できない
		fmt.Println(x, len(x), cap(x))
		fmt.Printf("%p \n", arr)
		fmt.Printf("%p \n", x)
	}
	// スライスをスライスすると「同じ場所を見てる」けど、一回appendすると別領域に差し替わるっぽい(仕様?)
	{
		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)
		y := x[:2:2]
		z := x[2:4:4]
		fmt.Println(cap(x), cap(y), cap(z))
		fmt.Printf("%p \n", x)
		fmt.Printf("%p \n", y)
		fmt.Printf("%p \n", &x[2])
		fmt.Printf("%p \n", z)
		y = append(y, 30, 40, 50)
		x = append(x, 60)
		z = append(z, 70)
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Printf("%p \n", x)
		fmt.Printf("%p \n", y)
		fmt.Printf("%p \n", z)
	}
	// 初手から「別メモリに領域確保する」んならcopy()使う
	{
		arr := []int{0, 1, 2, 3, 4, 5}
		arr_to := make([]int, len(arr)) // 第二引数で「copyする長さ」を指定しておかないと入っていかない(多分、メモリ的には入ってるんだろうけど使えない)
		copy_num := copy(arr_to, arr)   // copyは、恐らく「arr_toの内部のlenを変更しない」から、事前に指定しておかないと駄目(なので上の行で指定している)
		fmt.Println(arr)
		fmt.Printf("%p \n", arr)
		fmt.Println(arr_to, copy_num)
		fmt.Printf("%p \n", arr_to)
	}
	// 文字のスライス
	{
		s := "abcあいうdef"
		// byteでスライスすると「1バイトづつ」になる
		bs := []byte(s)
		// runeでスライスすると「1文字づつ」になる
		rs := []rune(s)
		//
		fmt.Println(bs)
		fmt.Println(rs)
	}
}

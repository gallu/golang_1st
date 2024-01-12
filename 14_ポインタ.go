package main

import "fmt"

type Hoge struct {
	s string
	i int
}

// ポインタじゃなく実体を渡してみる
func testSlice(arr []int) {
	fmt.Printf("\t%p \n", arr)
	arr[1] = 222
	fmt.Println("\t", arr, len(arr), cap(arr))
	fmt.Printf("\t%p \n", arr)
}

func printHoge(h Hoge) {
	fmt.Println("\t", h)
}

func dFunc1() int {
	fmt.Printf("dFunc1")
	return 1
}
func dFunc2() int {
	fmt.Printf("dFunc1")
	return 2
}
func dFunc3() int {
	fmt.Printf("dFunc1")
	return 3
}

func main() {
	// XXX 前提として「引数は値渡し(なので、いってもポインタの値渡しまで)なので以下略
	// スライスのポインタ周り
	{
		arr := []int{1, 2, 3, 4, 5}
		fmt.Println(arr, len(arr), cap(arr))
		fmt.Printf("%p \n", arr) // ようは配列だから、Cと一緒で「変数名がポインタ」
		testSlice(arr)
		// testSlice() の処理にもよるけど「もしかしたら同じポインタ見てるかもしれない」ので注意
		fmt.Println(arr, len(arr), cap(arr))
		fmt.Printf("%p \n", arr)
		// XXX mapも同じなので注意
	}

	// 構造体まわり
	{
		h := Hoge{s: "str", i: 100}
		ph := new(Hoge) // newは「その型のポインタ」を返す。中味は一通りゼロ値

		// とりあえず「型が違う」
		fmt.Println(h)
		fmt.Println(ph)

		// こっちは、まぁ
		h.i = 999
		// 「構造体ポインタ」の値アクセスもこれでよいっぽい
		ph.s = "hogera"
		ph.i = 999

		// 表示
		fmt.Println(h)
		fmt.Println(ph)

		printHoge(h)
		// printHoge(ph) // これは流石にアウト
		printHoge(*ph) // 実体化、これでよいぽい
	}

	// 関数ポインタ
	{
		// dfunc := dFunc1
		var dfunc func() int = dFunc1
		fmt.Printf("dfunc is %T \n", dfunc)
		i := (*&dfunc)()
		fmt.Println("i is ", i)
	}
}

package main

import "fmt"

// ポインタじゃなく実体を渡してみる
func testSlice(arr []int) {
	fmt.Printf("\t%p \n", arr)
	arr[1] = 222
	fmt.Println("\t", arr, len(arr), cap(arr))
	fmt.Printf("\t%p \n", arr)
}

func main() {
	// XXX 前提として「引数は値渡し(なので、いってもポインタの値渡しまで)なので以下略
	// スライスのポインタ周り
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Printf("%p \n", arr) // ようは配列だから、Cと一緒で「変数名がポインタ」
	testSlice(arr)
	// testSlice() の処理にもよるけど「もしかしたら同じポインタ見てるかもしれない」ので注意
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Printf("%p \n", arr)

	// XXX mapも同じなので注意
}

package pkg

import "fmt"

// いわゆるコンストラクタに近い感じ?
func init() {
	fmt.Println("pkg init!!")
}

// お外からは見れない(先頭1文字目が大文字じゃないから)
func hoge() {
	fmt.Println("hoge")
}

// お外から使える(先頭1文字目が大文字だから)
func Foo() {
	hoge() // 中からはcallできる
	fmt.Println("Foo")
}

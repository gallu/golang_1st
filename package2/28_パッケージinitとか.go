package main

//import pkg "my/pkg"
import "my/pkg"

func main() {
	// pkg.hoge() // 先頭小文字だから「お外からは使えない」
	pkg.Foo()
}

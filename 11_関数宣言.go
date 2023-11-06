package main

import "fmt"

// 可変長引数
func variableLengthFunc(key string, vals ...string) string {
	fmt.Println("variableLengthFunc:key", key)
	fmt.Println("variableLengthFunc:vals", vals)

	return "ok"
}

// 複数の戻り値
func returnMulti() (int, int, string) {
	return 10, 20, "ok"
}

// 名前付き戻り値
func namedReturn() (result int, err error) {
	result = 123
	//return	// これで通る。とはいえ、ブランクreturn(って言うっぽい)は非推奨っぽ
	return result, err // こっちのほうがいいぽい。「関数のなかで変数つかってない」ならゼロ値でreturnされる
}

func main() {
	// XXX 普通の関数呼び出しは省略
	// 「名前付き引数」とか「オプション引数」とかは「構造体でやってくれ」らしい

	// 可変長引数
	variableLengthFunc("mono")
	variableLengthFunc("a", "bbb", "ccc", "ddd")

	// 複数の戻り値
	{
		i, j, s := returnMulti()
		fmt.Println(i, j, s)
	}

	// 名前付き戻り値
	{
		i, _ := namedReturn()
		fmt.Println(i)
	}
}

package main

import (
	"errors"
	"fmt"
)

func test() (int, error) {
	return 0, errors.New("えららら")
}

// コレ使うとフォーマットできる
func test2() (int, error) {
	i := 123
	return 0, fmt.Errorf("%d でフォーマット指定とか出来る系 えらら", i)
}

// センチネルエラー用の定義をざっくり
var UserNotFoundError = errors.New("hogera not find")

func test3() (int, error) {
	return 0, UserNotFoundError
}

func main() {
	// アバウトにエラー発生させてキャッチしてみる
	{
		_, err := test()
		if err != nil {
			fmt.Println(err)
		}
	}
	{
		_, err := test2()
		if err != nil {
			fmt.Println(err)
		}
	}
	// センチネルエラー、といふものありけり
	{
		switch _, err := test3(); {
		case err == UserNotFoundError:
			fmt.Println("せんちねる 成功")
		case err != nil:
			fmt.Println(err)
		default:
			fmt.Println("エラーなかったから実行できたお!!")
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type Hoge struct {
	Num *int
	Str *string
	Fl  *float64
}

const s string = "const string"

func main() {
	// 値を使うとき 実態参照演算子必須 なの面倒だなぁ……
	{
		h := Hoge{}
		fmt.Printf("%#+v\n", h)

		// i := 100
		// h.Num = &i
		h.Num = new(int)
		*h.Num = 100
		h.Str = new(string)
		*h.Str = "Hogera mugera"
		fmt.Printf("%#v, %d\n", h, *h.Num)

		// とはいえ、jsonにするとちゃんと「値」になるんだなぁ
		s, err := json.Marshal(h)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s))
	}

	// 初期で値を入れておくと、まぁ……
	{
		num := 999
		types := s // XXX 定数直ポインタ使い、はNGっぽいので、一度うけとらないといけない……
		h := Hoge{
			Num: &num,
			Str: &types,
		}
		fmt.Printf("%#+v\n", h)

		//
		s, err := json.Marshal(h)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s))
	}

}

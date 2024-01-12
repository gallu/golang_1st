package main

import (
	"fmt"
	"strings"
)

func Snake2UpperCamel(base string) string {
	r := ""
	count := 0
	for _, v := range base {
		// アンダースコアなら(カウンタリセットして)すっ飛ばす
		if string(v) == "_" {
			count = 0
			continue
		}
		var add string
		if count == 0 {
			// 先頭なら大文字変換
			add = strings.ToUpper(string(v))
		} else {
			add = string(v)
		}
		r = r + add
		count++
	}
	return r
}

func main() {
	// 文字列の分解と結合
	{
		s := "1,2,3,4,5"
		awk := strings.Split(s, ",")
		fmt.Println(awk)

		s2 := strings.Join(awk, "_")
		fmt.Println(s2)
	}
	// 蛇をアッパーキャメルに
	{
		tests := []string{
			"abc_def_gij",
		}

		for _, v := range tests {
			r := Snake2UpperCamel(v)
			fmt.Println(v, "->", r)
		}
	}
}

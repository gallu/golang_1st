package main

import "fmt"

func main() {
	// いわゆるhash配列
	{
		// 空のmapの宣言
		harr := map[string]int{}
		fmt.Println(harr)
	}
	{
		harr := map[string]int{
			"key":  10,
			"hoge": 999,
		}
		fmt.Println(harr)
		// 要素の追加
		harr["foo"] = 123
		fmt.Println(harr)
		// keyの有無はこんな風に。「カンマokイディオム」って言うぽい
		// 二番目の戻り値がboolで、そこで判断するぽい
		v, ok := harr["key"]
		fmt.Println(v, ok)
		v, ok = harr["dummy"]
		fmt.Println(v, ok)

		// 要素の削除
		delete(harr, "key")
		v, ok = harr["key"]
		fmt.Println("key delete", v, ok)
	}
	// 値が配列の時の書き方
	{
		harr := map[string][]int{
			"key":  []int{10, 20},
			"hoge": []int{111, 2222},
		}
		fmt.Println(harr)
	}
}

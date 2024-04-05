package main

import "fmt"

func main() {
	{
		sl := []int{0, 1, 2}
		fmt.Printf("%#v \n", sl)
		// 先頭の除去(全体的に半開区間(左閉右開区間)なんだねぇ)
		sl = sl[1:]
		fmt.Printf("%#v \n", sl)
	}
	{
		sl := []int{0, 1, 2}
		fmt.Printf("%#v \n", sl)
		// 真ん中抜くのは面倒だねぇ
		pos := 1
		sl = append(sl[:pos], sl[pos+1:]...)
		fmt.Printf("%#v \n", sl)
	}
	{
		sl := []int{0, 1, 2}
		fmt.Printf("%#v \n", sl)
		// 末尾の除去
		sl = sl[:2]
		fmt.Printf("%#v \n", sl)
	}
	{
		sl := []int{0, 1, 2}
		fmt.Printf("%#v \n", sl)
		// これで先頭？
		pos := 0
		sl = append(sl[:pos], sl[pos+1:]...)
		fmt.Printf("%#v \n", sl)
	}
	{
		sl := []int{0, 1, 2}
		fmt.Printf("%#v \n", sl)
		// これで末尾？
		pos := 2
		sl = append(sl[:pos], sl[pos+1:]...)
		fmt.Printf("%#v \n", sl)
	}
}

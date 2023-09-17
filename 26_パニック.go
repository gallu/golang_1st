package main

import (
	"fmt"
)

func div60() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("ぱにっく!!: ")
			fmt.Println(v)
		}
	}()
	// panic("やだぁ")
	// fmt.Println(10 / 0) // これだとコンパイル時にバレちゃう
	i := 0
	fmt.Println(10 / i)
}

func main() {
	div60()
}

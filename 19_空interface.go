package main

import "fmt"

func main() {
	// 自在武器過ぎて、気をつけないとかなり危なそう
	{
		var v interface{}
		v = 123
		fmt.Println(v)
		fmt.Printf("%T \n", v)

		v = "abc"
		fmt.Println(v)
		fmt.Printf("%T \n", v)

		v = true
		fmt.Println(v)
		fmt.Printf("%T \n", v)

	}
	// anyは1.18から、らしい
	{
		var v any
		v = 123
		fmt.Println(v)
		fmt.Printf("%T \n", v)

		v = "abc"
		fmt.Println(v)
		fmt.Printf("%T \n", v)
	}
}

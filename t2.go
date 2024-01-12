package main

import "fmt"

var base = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	{
		s := "abc"
		fmt.Println(len(s))
	}

	{
		s := []byte("abc")
		fmt.Println(len(s))
	}
	{
		s := []rune("abc")
		fmt.Println(len(s))
	}
	{
		fmt.Println(len(base))
	}

}

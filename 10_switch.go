package main

import "fmt"

func main() {
	// 普通の
	{
		for _, v := range []int{1, 2, 3, 4} {
			switch v {
			case 1, 2:
				fmt.Println(v)
			case 3:
				fmt.Println("---", v)
				fallthrough // 明示すると下に落とせる
			default:
				fmt.Println("any", v)
			}
		}
	}

	// ブランクswitch
	{
		for _, v := range []string{"a", "bc", "def", "hogera"} {
			switch len := len(v); {
			case len == 1:
				fmt.Println("len 1")
			case len <= 3:
				fmt.Println("len <= 3")
			case len == 2: // 順番的に通らない
				fmt.Println("len == 2")
			default:
				fmt.Println("len to long")
			}
		}
	}
}

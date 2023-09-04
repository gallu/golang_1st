package main

import "fmt"

// XXX これでも「ここだけ」で使うんならいいけど「外に出る」時は運用気をつけないと怖いなぁ。そも「外に出せるのか?」ってのはあるにしても
const (
	AAA = iota
	BBB
	CCC
	DDD
)

func main() {
	// 列挙型の表示
	fmt.Println(AAA)
	fmt.Println(BBB)
	fmt.Println(CCC)
}

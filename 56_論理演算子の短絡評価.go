package main

import "fmt"

func rTrue() bool {
	fmt.Println(true)
	return true
}

func rFalse() bool {
	fmt.Println(false)
	return false
}

func main() {
	// 普通に短絡評価するねぇ。仕様なんかなぁ？
	fmt.Printf("rTrue() || rFalse() return is %v \n", rTrue() || rFalse())
	fmt.Printf("rFalse() || rTrue() return is %v \n", rFalse() || rTrue())
	fmt.Printf("rTrue() && rFalse() return is %v \n", rTrue() && rFalse())
	fmt.Printf("rFalse() && rTrue() return is %v \n", rFalse() && rTrue())
}

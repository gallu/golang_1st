package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	ch := make(chan string, 1)
	chErr := make(chan string, 1)

	// 乱数で処理を変える
	if n, _ := rand.Int(rand.Reader, big.NewInt(2)); n.Int64() == 0 {
		ch <- "ok"
		chErr <- ""
	} else {
		ch <- ""
		chErr <- "えらららなのだ"
	}

	// 受信して確認
	ret := <-ch
	err := <-chErr

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}

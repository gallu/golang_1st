package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	// バッファリングって「送信できる数量」なんだ。バイト数とかじゃなくて
	{
		fmt.Println("バッファリングって「送信できる数量」なんだ。バイト数とかじゃなくて")
		ch := make(chan string, 2)
		num := 10
		go func() {
			defer close(ch)
			for i := 0; i < num; i++ {
				n, _ := rand.Int(rand.Reader, big.NewInt(999))
				fmt.Println("sending....", i, n)
				ch <- fmt.Sprintf("data %d(%d) だよ!!", i, n)
			}
		}()

		// ぐるっと取得
		for v := range ch {
			fmt.Printf("received value: %s\n", v)
			time.Sleep(200 * time.Millisecond) // 0.2秒スリープ(くらい入れるとバッファリングされている状況が把握しやすい)
		}
	}
}

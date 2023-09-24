package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(ctx context.Context, wg *sync.WaitGroup, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		// コンテキストの値の取得
		userID := ctx.Value("userID").(int)  // 値は any(interface{})なので、型変換
		token := ctx.Value("token").(string) // 値は any(interface{})なので、型変換
		hoge := ctx.Value("hoge")            // 存在しない値は? nilが返るぽ。なのでifで判定とかすればよし(これ1文にでけんもんかね?)
		if hoge != nil {
			hoge = hoge.(string)
		}
		fmt.Println("userID", userID)
		fmt.Println("token", token)
		fmt.Println("hoge", hoge)

	LOOP:
		for {
			num++
			select {
			case <-ctx.Done(): // doneチャネルがcloseされたらbreakが実行される
				// ctx.Err() で理由が取れるぽ
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					// キャンセルされていた場合
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					// タイムアウトだった場合
					fmt.Println("deadline")
				}
				break LOOP
			case out <- num: // キャンセルされてなければnumを送信
				fmt.Println("generator for start", num)
			}
			fmt.Println("generator for end", num)
		}
		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

// 「各ゴルーチン間の情報伝達用」……なるほど
// https://zenn.dev/hsaki/books/golang-context/viewer/definition
func main() {
	{
		// doneの伝搬
		ctx, cancel := context.WithCancel(context.Background())
		// 値を渡す(ゴルーチン側から受け取るのは難しいぽ?)
		ctx = context.WithValue(ctx, "userID", 123)
		ctx = context.WithValue(ctx, "token", "abcdefg")
		//
		wg := &sync.WaitGroup{}
		gen := generator(ctx, wg, 0)

		wg.Add(1)

		for i := 0; i < 5; i++ {
			fmt.Println(<-gen)
		}
		cancel() // 5回genを使ったら、doneチャネルをcloseしてキャンセルを実行: ここで「<-ctx.Done()」がcallされる状況を整えるんだ

		wg.Wait()
	}
}

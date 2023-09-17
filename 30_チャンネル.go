package main

import (
	"fmt"
	"sync"
)

func main() {
	// とりあえず送受信
	{
		// 「stringを送受信する」チャンネルの作成
		ch := make(chan string)
		// チャンネルにデータを送りつける
		// XXX ここ、goroutineにしないと「 all goroutines are asleep - deadlock!」って怒られる!!
		go func() {
			// 送信
			ch <- "data string"
			// close
			close(ch)
		}()
		// チャンネルからデータを受け取る
		s := <-ch
		// 表示
		fmt.Println(s)

		// 「無いのに受信」は「カンマOKイディオム」で取れるぽい
		s, ok := <-ch
		if !ok {
			fmt.Println("データなかったぽ")
		} else {
			fmt.Println("データあった? まじ?", s)
		}
	}

	// 複数回の送受信
	{
		fmt.Println("複数回の送受信")
		ch := make(chan string)
		num := 5
		for i := 0; i < num; i++ {
			go func(i int) {
				ch <- fmt.Sprintf("data %d だよ!!", i)
			}(i)
		}
		// 1回だけ受信(したらどうなる?): 1回だけの受信だったお
		//	s := <-ch
		//	fmt.Println(s)

		// これで全部取れる
		// XXX けど順番はばらんばらん
		for i := 0; i < num; i++ {
			s := <-ch
			fmt.Println(s)
		}
		close(ch) // XXX 下参照。close場所、多分ここ「ではない」っぽい
	}

	// 受信側もgoroutineにしてみる
	{
		fmt.Println("受信側もgoroutineにしてみる")
		ch := make(chan string)
		num := 5
		var wg sync.WaitGroup
		wg.Add(num * 2)
		for i := 0; i < num; i++ {
			go func(i int) {
				defer wg.Done()
				ch <- fmt.Sprintf("data %d だよ!!", i)
			}(i)
		}

		// これで全部取れる
		// XXX 順番も整った……やっぱりこっちが本道???
		for i := 0; i < num; i++ {
			go func() {
				defer wg.Done()
				s := <-ch
				fmt.Println(s)
			}()
		}
		wg.Wait() // 待ってから
		close(ch) // XXX 下参照。close場所、多分ここ「ではない」っぽい
	}

	// 受信を「回数不明」状態にしてみる
	{
		fmt.Println("受信を「回数不明」状態にしてみる")
		ch := make(chan string)
		go func(ch chan string) { // XXX ここの引数の書き方、 ch chan<- string って野もあるらしい。後でわかったら追記する: 「定義のしかたによって、チャネルを送信専用・受信専用にすることができます。」らしい!!
			for i := 0; i < 5; i++ {
				ch <- fmt.Sprintf("data %d だよ!!", i)
			}
			close(ch) // 閉じる: 送信側でちゃんと閉じてやらないと駄目ぽい
		}(ch)

		// これで全部取れる
		for s := range ch {
			fmt.Println(s)
		}
	}
}

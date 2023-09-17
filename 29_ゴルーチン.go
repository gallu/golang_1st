package main

import (
	"fmt"
	"sync"
	"time"
)

// なんの制御もない感じ
func hoge(s string) {
	fmt.Println(s, "start")
	time.Sleep(500 * time.Millisecond) // 0.5秒スリープ
	fmt.Println(s, "end")
}

// 共有メモリにカウンタがあってwaitできる感じ
func foo(counter *sync.WaitGroup, s string) {
	// counter.Add(1)       // こっちのほうがあつかいやすくない? って思ったんだけど、元のほうで「ストンと下に抜けて終了しちゃう」から駄目なのか
	defer counter.Done() // カウンターのデクリメント。defer で「抜ける直前に処理」させる

	fmt.Println(s, "start")
	time.Sleep(500 * time.Millisecond) // 0.5秒スリープ
	fmt.Println(s, "end")
}

// mutex ロックあり
func bar(mu *sync.Mutex, counter *sync.WaitGroup, s string) {
	defer counter.Done() // カウンターのデクリメント。defer で「抜ける直前に処理」させる

	mu.Lock() // ロック確保

	fmt.Println(s, "start")
	time.Sleep(500 * time.Millisecond) // 0.5秒スリープ
	fmt.Println(s, "end")

	mu.Unlock() // ロック解放: これ、deferでもよくね?(処理全体にかかるんなら)
}

// 単純な並行処理のテスト
func test1() {
	// これだと「A end」が厳密には出てくる保証がない(実際、動かして稀に出てこなかった)
	go hoge("A")
	hoge("B") // ここも goroutine にしちゃうと「一気に通り抜けて終わり」になっちゃうｗ

	// 終わりを告げる(けどAはまだ終わってない)
	fmt.Println("test1 end")
}

// 「待ち処理」を入れた並行処理
func test2() {
	var counter sync.WaitGroup
	counter.Add(1)
	go foo(&counter, "AA")
	counter.Add(1)
	go foo(&counter, "BB")

	// 終了をまって終わりを告げる
	counter.Wait()
	fmt.Println("test2 end")
}

// mutex ロック、あるんだ
func test3() {
	var counter sync.WaitGroup
	var mu sync.Mutex // いったんここ。ただ「このデータのmutexを確保したい」時は、「このデータ」の構造体に入れちゃうほうがよさそうっぽい
	counter.Add(1)
	go bar(&mu, &counter, "AA")
	counter.Add(1)
	go bar(&mu, &counter, "BB")

	// 終了をまって終わりを告げる
	counter.Wait()
	fmt.Println("test2 end")
}

func main() {
	//test1()
	//test2()
	test3()
}

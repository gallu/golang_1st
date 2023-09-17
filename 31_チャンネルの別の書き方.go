package main

import "fmt"

// チャネルに値を渡す
// 引数の書き方が微妙にわからんな……「定義のしかたによって、チャネルを送信専用・受信専用にすることができます。」らしい!!
func sendValue(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("No %d", i)
	}
	close(c)
}

// チャネルから値を受け取る
func receiveValue(c <-chan string) {
	for v := range c {
		fmt.Println("チャネルから受け取った値：", v)
	}
}

// どかどかと送信する
func test(s [][]int, ch chan<- int) {
	// チャネルを閉じる
	// XXX なるほど「送信し終わる」系ならこの方がわかりやすいのかも
	defer close(ch)

	for _, v := range s {
		sum := 0
		for _, v2 := range v {
			// チャネルに送信
			ch <- v2
			// 合算
			sum += v2
		}
		// チャネルに送信
		ch <- sum
	}
}

func main() {
	{
		fmt.Println("送受信")
		c := make(chan string)
		go sendValue(c)
		receiveValue(c)

		fmt.Println("処理を終了します。")
	}

	// どかどかと送信してみる
	{
		fmt.Println("どかどかと送信してみる")
		// チャネルの作成
		ch := make(chan int)

		s := [][]int{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
		}
		go test(s, ch)

		// チャネルが閉じられるまで受信
		for i := range ch {
			fmt.Println(i)
		}
	}
}

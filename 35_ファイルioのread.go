package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 多分「一番ラフな」read
	// ファイルまるっと、だから、メモリはそれなりに
	{
		bytes, err := os.ReadFile("./read.txt")
		if err != nil {
			// ざっくりエラーチェック
			panic(err)
		}
		fmt.Printf("%T \n", bytes) // []uint8 なのか!!
		fmt.Print(string(bytes))
	}
	// ファイルが無い時の挙動
	{
		_, err := os.ReadFile("./dummy")
		if err != nil {
			// ざっくりエラーチェック
			fmt.Println(err)
		}
	}

	// ある程度メモリを意識するときは多分こんな感じ?
	{
		f, err := os.Open("./read.txt")
		defer f.Close() // 「閉じる」はとっとと予約しておく
		if err != nil {
			// ざっくりエラーチェック
			panic(err)
		}
		//  バッファ用意して
		bytes := make([]byte, 3)
		// ループして読み込む
		for {
			count, err := f.Read(bytes)
			// readが0バイトなら終了: 多分真っ当な終わり
			if count == 0 {
				fmt.Println("\n0 end")
				break
			}
			// えら?
			if err != nil {
				// EOFだったので普通におわり
				if err == io.EOF {
					fmt.Println("EOF end")
					break
				}
				// EOF以外ならえら終了
				panic(err)
			}
			// 読めたぽ
			// XXX これ「バッファ全部使って良い」のか「読めた文字数だけ使う」のかが微妙なんだが -> うんなんか「以前に読み込んだ変なデータが残る」みたい、なので、ちゃんとスライスしませう
			fmt.Print(string(bytes[:count]))
		}
		fmt.Println("fin")
	}

}

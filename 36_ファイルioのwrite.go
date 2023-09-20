package main

import "os"

func main() {
	// ラフい書き方
	{
		err := os.WriteFile("./write_1.txt", []byte("hello\ngo\n"), 0644)
		if err != nil {
			panic(err)
		}
	}

	// いわゆる「上書き」モードでのファイル書き込み
	// XXX パーミッションど～なるんだろうねぇ? まぁ後で設定すりゃいいんだが
	{
		f, err := os.Create("./write_2.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// 引数「byteのスライス」なんだねぇ
		f.Write([]byte("hogera\n"))
	}

	// append(追加書き)は、こっちしかやり方ないんかねぇ?
	{
		// XXX 厳密には「os.O_CREATE いらん(から第三引数もいらん)」のだけど、学習コードなので付けるだけ付けておく
		f, err := os.OpenFile("./write_2.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// 引数「byteのスライス」なんだねぇ
		f.Write([]byte("fooobar\n"))
	}
}

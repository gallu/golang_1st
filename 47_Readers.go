package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// stringからのreaderの作成
	{
		// readerの作成
		r := strings.NewReader("hogera mugera foo bar.")
		// 読み込むバッファの作成
		b := make([]byte, 4)

		// ぶん回す
		for {
			// nは「読み込んだ文字数」ぽい
			// XXX バッファのスライスが2以上だと「スライスのcap()未満を読み込んだ」時に後ろにゴミが残るので注意!!
			n, err := r.Read(b)
			fmt.Printf("%v(%T)) %T, %v, %s \n", n, n, b, b, string(b))

			// 最後まで呼んだお!!
			// XXX readのnが0、って条件でもいけそうではあるんだが……
			if err == io.EOF {
				break
			}
		}

	}
}

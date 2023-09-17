package main

import (
	"errors"
	"fmt"
)

func main() {
	// エラーのラップ / unラップ
	{
		// こっちは普通のエラー
		err := errors.New("ほげら")
		fmt.Println(err)

		// ラッピング
		// XXX fmtは「どうなん?」って思うんだけど、現状こうなるっぽい……
		err2 := fmt.Errorf("ふげら。元々は %w", err)
		fmt.Println(err2)
		// 取り出す(unラップ)
		errBase := errors.Unwrap(err2)
		fmt.Println(errBase)
	}

	// エラーのスタックトレース(まだ標準では入ってないぽい)
	// XXX のでいったんオミット。後でやるかも(2.0で標準に入るかも? とか……入るといいなぁ……)
}

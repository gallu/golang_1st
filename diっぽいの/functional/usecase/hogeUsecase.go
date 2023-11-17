package usecase

import (
	"functional/repositories"
)

func Hoge(i int, additionArgs ...hogeOption) {
	// デフォルト値をここで設定
	args := &hogeOptions{}

	// デフォルト引数で「明示的に指定された値」を取得
	for _, setter := range additionArgs {
		setter(args)
	}
	// 設定がなかったらデフォ値設定
	if args.hr == nil {
		args.hr = repositories.NewHogeRepository()
	}

	// コンテナの値を使う
	args.hr.AnyDo(i)
}

// Functional Option Pattern用
// デフォルトで引数になりうる値の群れを構造体で宣言
type hogeOptions struct {
	hr repositories.HogeRepositorier
}

// 「デフォルトの引数の値の解決」をするための処理
type hogeOption func(*hogeOptions)

// これ、外部から使うから、publicで
// DIしたいのが２つ以上あるときは、この関数の引数を２つ以上にして「全部注入」「全部デフォ」のほうが多分とりまわしがいい
func ParamHogeRepository(hr repositories.HogeRepositorier) hogeOption {
	return func(args *hogeOptions) {
		args.hr = hr
	}
}

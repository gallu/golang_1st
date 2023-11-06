package usecase

import (
	"testinit/repositories"
)

// 本来ルート用のラッパー
func Hoge(i int) {
	// 引数を「本来のrepository」指定して本体をcall
	hoge(i, repositories.NewHogeRepository())
}

// テストルート用のラッパー
// XXX 多分これ、実際には何種類か作るんだと思う
func HogeTest(i int) {
	// 引数を「テスト用のrepository」指定して本体をcall
	hoge(i, repositories.NewHogeRepositoryMock())
}

// 本体
func hoge(i int, hr repositories.HogeRepositorier) {
	// コンテナの値を使う
	hr.AnyDo(i)
}

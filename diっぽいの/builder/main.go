package main

import (
	"fmt"

	"builder/repositories"
	"builder/usecase"
)

func main() {
	// 通常の呼び出し
	{
		fmt.Println("通常の呼び出し")
		uc := usecase.NewHogeUsecase()
		uc.Hoge(10)
	}
	// (主にテスト想定)DIで依存性ぶちこんでの呼び出し
	// XXX 最後 Build() とかのメソッドで〆てないからもどきだけど、Builder パターンっぽいナニカ
	{
		fmt.Println("テスト用の呼び出し")
		uc := usecase.NewHogeUsecase().
			HogeRepository(repositories.NewHogeRepositoryMock())
		uc.Hoge(10)
	}
}

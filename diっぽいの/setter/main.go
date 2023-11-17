package main

import (
	"fmt"
	"setter/repositories"
	"setter/usecase"
)

func main() {
	// 通常の呼び出し
	{
		fmt.Println("通常の呼び出し")
		uc := usecase.NewHogeUsecase()
		uc.Hoge(10)
	}
	// (主にテスト想定)DIで依存性ぶちこんでの呼び出し

	{
		fmt.Println("テスト用の呼び出し")
		uc := usecase.NewHogeUsecase()
		uc.Hr = repositories.NewHogeRepositoryMock()
		uc.Hoge(10)
	}
}

package usecase

import (
	"setter/repositories"
)

/*
 * これ、DI対象を外の変数で持つと「広域」になっちゃって影響がでかいので、構造体の中に閉じ込める(から少し作りを変える)
 */
type HogeUsecase struct {
	// いったん「直接外から触る」。setter経由のほうがよかったらそのように
	Hr repositories.HogeRepositorier
}

func NewHogeUsecase() *HogeUsecase {
	return &HogeUsecase{
		Hr: repositories.NewHogeRepository(),
	}
}

// 本体
func (h *HogeUsecase) Hoge(i int) {
	h.Hr.AnyDo(i)
}

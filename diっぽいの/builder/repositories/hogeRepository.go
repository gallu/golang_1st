package repositories

import (
	"fmt"
)

// 構造体用意
type HogeRepository struct{}

func NewHogeRepository() *HogeRepository {
	return new(HogeRepository)
}

func (hr *HogeRepository) AnyDo(i int) {
	fmt.Println("\tHogeRepository in func", i)
}

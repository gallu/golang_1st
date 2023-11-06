package repositories

import (
	"fmt"
)

// 構造体用意
type HogeRepositoryMock struct{}

func NewHogeRepositoryMock() *HogeRepositoryMock {
	return new(HogeRepositoryMock)
}
func (hr *HogeRepositoryMock) AnyDo(i int) {
	fmt.Println("\tHogeRepositoryMock in func", i)
}

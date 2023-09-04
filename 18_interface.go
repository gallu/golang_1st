package main

import "fmt"

// 武器
type Weapon struct {
	name    string
	ofPoint int
	price   int
}

// 鎧
type Armor struct {
	name     string
	defPoint int
	price    int
}

// 盾
type Shield struct {
	name     string
	defPoint int
	price    int
}

// interface
// XXX ようは「このメソッドもってたらこのグループに入れてあげるお!!」って事かし
type itemInterface interface {
	getPrice() int
}

// メソッド
func (w Weapon) getPrice() int {
	return w.price
}
func (a Armor) getPrice() int {
	return a.price
}

// もういっちょメソッド
// XXX ここでinterface使ってみる: 「このinterface持ってる構造体なら型とて許容する」って感じだよねぇ
func getPrice(item itemInterface) int {
	return item.getPrice()
}
func getSellPrice(item itemInterface) int {
	return item.getPrice() / 2
}

func main() {
	w := Weapon{
		name:    "ヒノキの棒",
		ofPoint: 1,
		price:   10,
	}
	a := Armor{
		name:     "ローブ",
		defPoint: 0,
		price:    20,
	}
	//
	fmt.Println(w.name)
	fmt.Println(getPrice(w))
	fmt.Println(getSellPrice(w))

	fmt.Println(a.name)
	fmt.Println(getPrice(a))
	fmt.Println(getSellPrice(a))

	// 盾はメソッド切ってないから使えない
	s := Shield{
		name:     "バックラー",
		defPoint: 1,
		price:    50,
	}
	fmt.Println(s.name)
	// fmt.Println(getPrice(s)) // ここが使えない: 型が不一致だから
}

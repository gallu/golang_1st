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

// Stringer というインタフェースがあるらしい
// https://go-tour-jp.appspot.com/methods/17
func (w Weapon) String() string {
	return fmt.Sprintf("武器) %s: %d / %d", w.name, w.ofPoint, w.price)
}
func (s Shield) String() string {
	return fmt.Sprintf("盾) %s: %d / %d", s.name, s.defPoint, s.price)
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
		name:     "いいぢす",
		defPoint: 1,
		price:    50,
	}
	fmt.Println(s.name)
	// fmt.Println(getPrice(s)) // ここが使えない: 型が不一致だから

	// Stringerに引っかかる(== String()が実装されている)と、そいつの結果を出す。そうじゃなきゃ中味を一通り。PHPの __toString() みたいな感じかしらん?
	fmt.Println(w)
	fmt.Println(a)
	fmt.Println(s)
}

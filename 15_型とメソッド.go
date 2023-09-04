package main

import "fmt"

// 構造体
type User struct {
	id        int
	lastName  string
	firstName string
	age       int
}

// 型メソッド(あるいはメソッド)
// 値レシーバー
func (u User) GetName() string {
	return u.lastName + " " + u.firstName
}

// ポインタレシーバー
func (u *User) AddAge() {
	// 書き方、デリファレンスいらないから、C++の & みたいなイメージだねぇ
	u.age++
}

//
func main() {
	// いわゆる クラス っぽい感じで作ってみる
	{
		u := User{
			id:        1,
			lastName:  "michi",
			firstName: "furu",
			age:       20,
		}
		name := u.GetName()
		fmt.Println(name)

		u.AddAge()
		fmt.Println(u.age)
		// 関数だからメソッド(関数ポインタ)もげとれる
		fn := u.AddAge
		fn()
		fmt.Println(u.age)
	}
}

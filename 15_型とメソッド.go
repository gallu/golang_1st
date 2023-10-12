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

// てけとうな型
type Hint int

// 「てけとうな型」のレシーバ
func (hi Hint) add(i int) Hint {
	return hi + Hint(i) // キャストしないと駄目
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

		// レシーバを呼ぶ時、構造体の変数が「実体」でも「ポインタ型」でも、特に問題なし(中でうまいことやってくれてるんでしょう)
		up := &u
		fmt.Printf("u is %T, up is %T \n", u, up)
		fmt.Println(u)
		u.AddAge()
		fmt.Println(up.age) // あえて「ポインタ型」のほうを使ってみる
		up.AddAge()
		fmt.Println(u.age) // あえて実体のほうを使ってみる
		(*up).AddAge()     // もちろんこの書き方もある(多分「やらない」んだろうけど)
		fmt.Println(up.age)
	}

	// 「てけとうな型」にもレシーバ(メソッド)が付与できる
	{
		var hi Hint = Hint(10)
		fmt.Println(hi)
		hi = hi.add(20)
		fmt.Println(hi)
	}
}

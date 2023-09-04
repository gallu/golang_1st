package main

import "fmt"

type User struct {
	id   int
	name string
}

// 継承(埋め込み)
type PlayUser struct {
	User
	gameName string
}

// なんか慣習的にこーゆーメソッド作る事も多い、の???
func newPlayUser(id int, name, gameName string) *PlayUser {
	p := new(PlayUser)
	p.id = id
	p.name = name
	p.gameName = gameName

	return p
}

// メソッド
func (u User) getName() string {
	return u.name
}

// PCのメソッド
func (p PlayUser) getData() string {
	return p.name + " " + p.gameName
}

func main() {
	{
		p := PlayUser{
			// 明示しなきゃいけないんだ
			User: User{
				id:   1,
				name: "gallu",
			},
			gameName: "TRPG",
		}
		fmt.Println(p)
		fmt.Println(p.getName()) // これは通るんだ
		fmt.Println(p.getData())
	}
	// こーゆー作り方もできるぽい
	{
		// newで「その構造体のポインタ型」がげとれるぽい
		pp := new(PlayUser)
		fmt.Printf("%T, %p \n", pp, pp)
		pp.id = 2
		pp.name = "mel"
		pp.gameName = "GURPS"
		fmt.Println(pp)
		fmt.Println(pp.getName())
		fmt.Println(pp.getData())
	}
	// newっぽいこと
	{
		p := newPlayUser(1, "may", "DnD")
		fmt.Println(p)
		fmt.Println(p.getName())
		fmt.Println(p.getData())
	}
}

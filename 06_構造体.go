package main

import (
	"fmt"
)

type user struct {
	id   int
	name string
}

func main() {
	{
		u := user{
			id:   1,
			name: "hogera",
		}
		fmt.Println(u)
		fmt.Printf("%T \n", u) // 型名の確認
	}
	// typeで型宣言しない無名の構造体も作れる
	{
		u := struct {
			id   int
			name string
		}{
			id:   1,
			name: "hohoho",
		}
		fmt.Println(u)
		fmt.Printf("%T \n", u) // 型名の確認
	}
	// 内部にも定義できる(scope内だけで有効)
	{
		type user_type_2 struct {
			id   int
			name string
		}
		u := user_type_2{
			id:   1,
			name: "hogera",
		}
		// 「全く同じフォーマット」なら変換できる
		var uu user = user(u)
		fmt.Println(uu)
		fmt.Printf("%T \n", uu) // 型名の確認
	}

	/*
		// 「順番違う」でもNG
		{
			type user_type_3 struct {
				name string
				id   int
			}
			u := user_type_3{
				id:   1,
				name: "hogera",
			}
			var uu user = user(u)
			fmt.Println(uu)
			fmt.Printf("%T \n", uu) // 型名の確認
		}
	*/

	// 構造体の比較
	{
		/*
			P 113
			(片方が無名構造体なら)keyがあってればいい? 順番は?
		*/
		// 「中味がいっしょ」ならtrueっぽい
		u := user{
			id:   1,
			name: "hogera",
		}
		u2 := user{
			id:   1,
			name: "hogera",
		}
		fmt.Printf("%p \n", &u)
		fmt.Printf("%p \n", &u2)
		fmt.Println(u == u2)
		u2.id = 2
		fmt.Println(u == u2)

		/*
			type user_type_2 struct {
				id   int
				name string
			}
			uType2 := user_type_2{
				id:   1,
				name: "hogera",
			}
			// これ、そもそも型が違うから「invalid operation」になる
			// 「代入可能」なものでもアウトなんだ
			fmt.Println(u == uType2)
		*/

		// 無名構造体は「順番含めてカラム名と中味が一致してたら」true
		uUnknown := struct {
			id   int
			name string
		}{
			id:   1,
			name: "hogera",
		}
		fmt.Println(u == uUnknown)
		uUnknown.id = 2
		fmt.Println(u == uUnknown)
		/*
			// 無名構造体、カラム順番が違うと「invalid operation」
			uUnknown2 := struct {
				name string
				id   int
			}{
				id:   1,
				name: "hogera",
			}
			fmt.Println(u == uUnknown2)
		*/
	}

	// 構造体とポインタ
	{
		u := &user{
			id:   1,
			name: "hogera",
		}
		fmt.Printf("型はポインタ %T \n", u)

		// 基本は「実体参照」してからアクセス
		(*u).id = 777
		fmt.Println(u)

		// でも実体参照、省略できるぽ
		// XXX この辺微妙にC++と違うから注意
		u.id = 999
		fmt.Println(u)
	}

	// 「空」の確認
	{
		u := user{}
		fmt.Println(u)
	}
}

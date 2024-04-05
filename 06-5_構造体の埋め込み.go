package main

import (
	"encoding/json"
	"fmt"
)

type HogeBase struct {
	Id   int
	Name string
}

type Foo struct {
	HogeBase
	Age int
}

type Bar struct {
	HogeBase
	Id    int
	Name2 string
}

func main() {
	// まず使う
	{
		f := &Foo{
			HogeBase: HogeBase{
				Id:   1,
				Name: "name",
			},
			Age: 123,
		}
		fmt.Printf("%#v\n", f)
		fmt.Println(f.Id, f.Name)

		s, err := json.Marshal(f)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s)) // {"Id":1,"Name":"name","Age":123}
	}

	// 「空からえっちらおっちら」
	{
		f := &Foo{}
		f.Id = 2
		f.Name = "foo name"
		f.Age = 82

		fmt.Printf("%#v\n", f)
		fmt.Println(f.Id, f.Name)

		s, err := json.Marshal(f)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s)) // {"Id":1,"Name":"name","Age":123}
	}

	// 埋め込んだ構造体へのアクセスいろいろ
	{
		f := &Foo{
			HogeBase: HogeBase{
				Id:   1,
				Name: "name",
			},
			Age: 123,
		}
		// ダイレクトでも埋め込み名はさんでもどっちでもいける
		fmt.Println(f.Id, f.HogeBase.Id)

		b := &Bar{
			HogeBase: HogeBase{
				Id:   1,
				Name: "hoge",
			},
			Id:    123,
			Name2: "bar name",
		}
		// 名前が競合すると「埋め込んでない」ほうが優先っぽい(仕様だと思うんだが、どうなんかねぇ?)
		fmt.Println(b.Id, b.Name)
		// 指定すればそれは当然「埋め込んだ側」が使われる
		fmt.Println(b.HogeBase.Id, b.HogeBase.Name)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Foo  []*FooMono
	Hoge struct {
		Id   int
		Name string
	}
}

type FooMono struct {
	Id   int
	Name string
}

type M map[string]interface{}

func main() {
	// めんどい……
	{
		d := Data{
			Hoge: struct {
				Id   int
				Name string
			}{
				Id:   1,
				Name: "hoge name",
			},
			Foo: []*FooMono{&FooMono{Id: 111, Name: "Foo name"}},
		}
		// json化
		json_string, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json_string))
	}
	// 面白い、けど、わかりやすいか?
	// 「がっつり動的」ならともかく「一定のフォーマットは死守したい」だとちょいと厳しいか?
	{
		d := M{
			"Hoge": M{
				"Id":   1,
				"Name": "hoge name",
			},
			"Foo": []*FooMono{&FooMono{Id: 111, Name: "Foo name"}},
		}
		// json化
		json_string, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json_string))
	}
}

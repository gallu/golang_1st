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
}

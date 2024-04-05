package main

import "fmt"

type Base struct {
	Id   int
	Name string
}

type Data struct {
	HogeID int
	Hoge   *Base

	FooID int
	Foo   *Base
}

func main() {
	// これはまぁできる(当たり前)
	{
		d := Data{
			HogeID: 1,
			Hoge: &Base{
				Id:   1,
				Name: "hoge",
			},
			FooID: 2,
			Foo: &Base{
				Id:   2,
				Name: "foo",
			},
		}
		fmt.Printf("%+v\n", d)
		fmt.Printf("%+v\n", d.Hoge)
		fmt.Printf("%+v\n", d.Foo)
	}

	// あ、やっぱりできるんだうん便利
	{
		d := Data{}
		{
			id := &d.HogeID
			(*id) = 1
			col := &d.Hoge
			(*col) = &Base{
				Id:   1,
				Name: "hoge",
			}
		}
		{
			id := &d.FooID
			(*id) = 2
			col := &d.Foo
			(*col) = &Base{
				Id:   2,
				Name: "foo",
			}
		}
		fmt.Printf("%+v\n", d)
		fmt.Printf("%+v\n", d.Hoge)
		fmt.Printf("%+v\n", d.Foo)
	}
}

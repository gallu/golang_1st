package main

import (
	"encoding/json"
	"fmt"
)

type Hoge struct {
	Id   int
	Name string
}

type Foo struct {
	Id int
}

type Bar struct {
	Id   int
	Name string
	F    Foo
}

type Fuga struct {
	Id  int
	Arr []int
}

type Piyo struct {
	Id int
	Mp map[int]string
}

func main() {
	{
		h := &Hoge{
			Id:   1,
			Name: "hoge",
		}
		// ポインタなので「同じもの」扱い
		h2 := h
		fmt.Printf("h : %p, %+v\n", h, h)
		fmt.Printf("h2: %p, %+v\n", h2, h2)

		// 実態参照しているので「別物」になる
		h3 := (*h)
		fmt.Printf("h : %p, %+v\n", h, h)
		fmt.Printf("h3: %p, %+v\n", &h3, h3)

		h.Id = 123
		h.Name = "changed hoge"
		fmt.Printf("h : %p, %+v\n", h, h)
		fmt.Printf("h2: %p, %+v\n", h2, h2)
		fmt.Printf("h3: %p, %+v\n", &h3, h3)

	}
	// 中に構造体があっても基本の動きは一緒だねぇ
	// ポインタじゃないから「別物」になる
	{
		h := Bar{
			Id:   1,
			Name: "bar",
			F: Foo{
				Id: 111,
			},
		}
		h2 := h
		fmt.Printf("h : %p, %+v\n", &h, h)
		fmt.Printf("h2: %p, %+v\n", &h2, h2)

		//
		h.Id = 123
		h.Name = "changed bar"
		h.F.Id = 999
		fmt.Printf("h : %p, %+v\n", &h, h)
		fmt.Printf("h2: %p, %+v\n", &h2, h2)
	}

	// スライスがあるやつのcopy
	// スライスまでなら実態の := でcopyでけるぽ
	{
		f := Fuga{
			Id:  1,
			Arr: []int{1, 2, 3},
		}
		// fmt.Printf("%+v\n", f)
		f2 := f
		f.Id = 123
		f.Arr = append(f.Arr, 1)

		fmt.Printf("%p, %+v\n", &f, f)
		fmt.Printf("%p, %+v\n", &f2, f2)
	}

	// map以外はcopyでけるんだけど、mapだけダメぽ
	{
		p := Piyo{
			Id: 1,
			Mp: make(map[int]string),
		}
		p2 := p
		p.Id = 123
		p.Mp[1] = "hoge"
		p.Mp[2] = "foo"
		fmt.Printf("%+v\n", p)
		fmt.Printf("%+v\n", p2)
	}

	// 「jsonにして戻す」とcopyはでける
	{
		p := Piyo{
			Id: 1,
			Mp: make(map[int]string),
		}
		p2 := &Piyo{}
		b, _ := json.Marshal(p)
		// err := json.Unmarshal(b, p2)
		json.Unmarshal(b, p2)

		p.Id = 123
		p.Mp[1] = "hoge"
		p.Mp[2] = "foo"
		fmt.Printf("%+v\n", p)
		fmt.Printf("%+v\n", *p2)

	}

}

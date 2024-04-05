package main

import "fmt"

type TypeHoge int

const (
	Hoge TypeHoge = 1
	Foo  TypeHoge = 2
	Bar  TypeHoge = 3
)

func (t TypeHoge) toString() string {
	r := ""
	switch t {
	case 1:
		r = "hoge"
	case 2:
		r = "foo"
	case 3:
		r = "bar"
	}
	return r
}

func main() {
	{
		v := Hoge
		fmt.Println(v, v.toString())
	}
}

package main

import (
	"fmt"
	"reflect"
)

type Hoge struct {
	id int
}

func (h Hoge) Func() {
}
func GetName(h HogeI) string {
	rv := reflect.ValueOf(h)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv.Type().Name()
}

type HogeI interface {
	Func()
	// GetName() string
}

func NewHoge(i int) HogeI {
	return &Hoge{
		id: i,
	}
}

func main() {
	{
		h := NewHoge(1)
		// fmt.Printf("%#v\n", h)
		fmt.Println(GetName(h))
	}
	{
		var h HogeI = Hoge{
			id: 1,
		}
		fmt.Println(GetName(h))
	}
	/*
		// これ(ダブルポインタをインタフェースで扱う)はできないっぽい
		{
			var h = &Hoge{
				id: 1,
			}
			var hh HogeI = &h
			fmt.Println(GetName(hh))
		}
	*/

	/*
		rf := reflect.TypeOf(h)
		fmt.Println("TypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
		rv := reflect.ValueOf(h)
		fmt.Println("ValueOf 型情報(Kind):", rv.Kind(), ", 型(Type):", rv.Type(), ", 値(Interface):", rv.Interface())

		if rf.Kind() == reflect.Ptr { // 多分「こっち」なんだろうなぁ
			rf := reflect.TypeOf(h).Elem()
			fmt.Println("\tTypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
			rv := reflect.ValueOf(h).Elem()
			fmt.Println("\tValueOf 型情報(Kind):", rv.Kind(), ", 型(Type):", rv.Type(), ", 値(Interface):", rv.Interface())
		}
		fmt.Println("")
	*/
}

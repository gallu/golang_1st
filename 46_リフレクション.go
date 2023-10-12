package main

import (
	"fmt"
	"reflect"
)

type IntSlice []int

type HogeStruct struct {
	id int
}

// 確認と表示
// いくつか確認してみる。これ以外にもかなり色々あるぽい
func p(v any) {
	rf := reflect.TypeOf(v)
	fmt.Println("TypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
	rv := reflect.ValueOf(v)
	fmt.Println("ValueOf 型情報(Kind):", rv.Kind(), ", 型(Type):", rv.Type(), ", 値(Interface):", rv.Interface())

	// ポインタなら中を掘ってみる
	//if rf.Kind().String() == "ptr" { // 文字列比較でもいいけど……
	if rf.Kind() == reflect.Ptr { // 多分「こっち」なんだろうなぁ
		rf := reflect.TypeOf(v).Elem()
		fmt.Println("\tTypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
		rv := reflect.ValueOf(v).Elem()
		fmt.Println("\tValueOf 型情報(Kind):", rv.Kind(), ", 型(Type):", rv.Type(), ", 値(Interface):", rv.Interface())
	}
	fmt.Println("")
}

func main() {
	//
	// 色々打ち込んでみる
	var s string = "abc"
	p(s)

	var i int = 123
	p(i)

	var i64 int64 = 123
	p(i64)

	// 配列とスライス
	arr := [3]int{1, 2, 3}
	p(arr)
	p(arr[0])
	sarr := []int{1, 2, 3}
	p(sarr)
	p(sarr[0])

	// typeで別名作ったら?
	tarr := IntSlice{1, 2, 3}
	p(tarr)

	// map
	m := map[string]int{
		"hoge": 10,
		"foo":  20,
	}
	p(m)

	// makeは?
	ms := make([]int, 10)
	p(ms)

	// 構造体
	hs := HogeStruct{
		id: 10,
	}
	p(hs)

	// new
	nhs := new(HogeStruct)
	p(nhs)

	// ポインタ
	ps := &s
	p(ps)

	// 関数もいける???
	p(p)

}

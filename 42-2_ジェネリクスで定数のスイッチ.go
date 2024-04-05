package main

import (
	"fmt"
	"reflect"
)

type CHoge uint64
type CFoo uint64

const (
	H1 CHoge = iota + 1
	H2
	H3
)
const (
	F1 CFoo = iota + 10
	F2
	F3
)

func Test[T CHoge | CFoo](p T) {
	// fmt.Printf("%T, %#v \n", p, p)
	/*
		// リフレクション使ってもいいんだけど、ゴツくない？
		rf := reflect.TypeOf(p)
		fmt.Println("TypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
		rv := reflect.ValueOf(p)
		fmt.Println("ValueOf 型情報(Kind):", rv.Kind(), ", 型(Type):", rv.Type(), ", 値(Interface):", rv.Interface())
	*/
	// tp := reflect.TypeOf(p).Name()
	tp := reflect.ValueOf(p).Type()
	fmt.Println(tp)
}

func main() {
	Test(H1)
	Test(F2)
}

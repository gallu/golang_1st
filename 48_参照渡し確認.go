package main

import (
	"fmt"
	"reflect"
)

func Test[T any](v, d T) {
	v = d
}

func Tarr(arr [2]int) {
	fmt.Printf("func\t%p \n", &arr)
}

func TSlice(sl []int) {
	fmt.Printf("func\t%p \n", sl)
	sl[2] = 999
	sl[3] = 999
}

func TMap(h map[string]int) {
	fmt.Printf("func\t%p \n", h)
	h["hoge"] = 999
}

func TStruct(sh SHoge) {
	// fmt.Printf("func\t%p \n", sh) // 怒られる
	sh.id = 999
	sh.name = "gallu"
}

type SHoge struct {
	id   int
	name string
}

func main() {
	// 実態で渡る系
	// 基本型は一通りアウトそう
	i := 1
	// fmt.Printf("%p \n", i) // 怒られる
	Test[int](i, 999)
	fmt.Println(i)

	b := true
	// fmt.Printf("%p \n", b) // 怒られる
	Test[bool](b, false)
	fmt.Println(b)

	s := "str"
	// fmt.Printf("%p \n", s) // 怒られる
	Test[string](s, "hogera")
	fmt.Println(s)

	f := 3.141592
	// fmt.Printf("%p \n", f) // 怒られる
	Test[float64](f, 1.41421356)
	fmt.Println(f)

	arr := [2]int{1, 2}
	fmt.Printf("main\t%p \n", &arr)
	Tarr(arr) // 配列だから加減算はできない(からまぁテストする意味もあんまりないんだけど)

	sh := SHoge{}
	fmt.Println(sh)
	// fmt.Printf("main\t%p \n", sh) // 怒られる
	TStruct(sh)
	fmt.Println(sh)

	// ポインタで渡る系(どうも、スライスとmapがポインタ渡しになるぽ)
	sl := make([]int, 10)
	fmt.Println(sl)
	fmt.Printf("main\t%p \n", sl)
	TSlice(sl)
	fmt.Println(sl)
	// リフレクションで確認してみる。型は(ptrじゃなくて)sliceになる。からそもそもとして「sliceはポインタ」ってのが前提なんだと思う
	rf := reflect.TypeOf(sl)
	fmt.Println("TypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())

	h := make(map[string]int)
	fmt.Println(h)
	fmt.Printf("main\t%p \n", h)
	TMap(h)
	fmt.Println(h)
	// リフレクションで確認してみる。型は(ptrじゃなくて)mapになる。からそもそもとして「mapはポインタ」ってのが前提なんだと思う
	rf = reflect.TypeOf(h)
	fmt.Println("TypeOf  型情報(Kind):", rf.Kind(), ", 型(Name):", rf.Name(), ", サイズ(Size):", rf.Size())
}

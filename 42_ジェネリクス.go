package main

import "fmt"

// べたっと書いてみる
func test[T int | float64](a, b T) T {
	return a + b
}

// 型をinterfaceで定義
type Hoge interface {
	int | float64
}

func test2[T Hoge](a, b T) T {
	return a + b
}

// チルダ(underlying type)を確認
type MyInt int
type MyFloat float64
type Foo interface {
	~int | ~float64
}

func test3[T Foo](a, b T) T {
	return a + b
}

// スライス確認
func test4[T []int](arr T) T {
	ret := append(arr, 999)
	return ret
}

func main() {
	// ベタっと
	{
		i := test(1, 2)
		fmt.Println(i)
		f := test(1.1, 2.2)
		fmt.Println(f)
		/*
			// これはやっぱり駄目だねぇ
			m := test(1.1, 2)
			fmt.Println(m)
		*/
	}
	// interfaceで
	{
		i := test2(1, 2)
		fmt.Println(i)
		f := test2(1.1, 2.2)
		fmt.Println(f)
	}
	// チルダ(underlying)の確認
	{
		var a MyInt = 10
		var b MyInt = 20
		// i := test(a, b) // コレだと「型エラー」になる
		i := test3(a, b) // こっちだと「MyIntは結局intだから」OKになる
		fmt.Println(i)
	}
	// スライス使える?
	{
		arr := []int{1, 2, 3}
		arr = test4(arr)
		fmt.Println(arr)
	}
}

package main

import "fmt"

// defer用関数
func doDefer(s string) {
	fmt.Println("defer", s)
}

func test() {
	fmt.Println("in test trap 1")

	// 無名関数でdefer
	// XXX scopeが外れるタイミングでうごくぽい
	defer func() {
		fmt.Println("in test derer")
	}()

	fmt.Println("in test end")
}

func main() {
	// ブロックは関係ないぽい(関数単位とかその辺?)
	{
		fmt.Println("in main trap 1")

		// LIFO
		// ブロックは関係ないぽい(関数単位とかその辺?)
		defer doDefer("in main 1")
		defer doDefer("in main 2")
		defer doDefer("in main 3")

		test()
	}
	fmt.Println("in main trap 1")

	// 値の評価タイミング
	i := 1
	defer func() {
		// XXX 実際に「そのscopeでdeferな関数がcallされるタイミングの値」っぽい、ので、今回だと999になる
		fmt.Println("in defer i is ", i)
	}()
	i = 999
}

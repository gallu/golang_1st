package main

import (
	"fmt"
	// モジュールからのパスを記述
	mypkg "myPkg/mypkgDir"
	"myPkg/nest/pkg1"
)

func main() {
	fmt.Println("Hello World!")
	mypkg.HelloFromMypkg()
	mypkg.Hello2()
	pkg1.TestFile()
	testTest()
}

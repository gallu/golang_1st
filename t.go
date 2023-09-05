/*
"Accept Interfaces, Return structs" 用のサンプルコード。まだ理解できないので後で咀嚼し直す
*/

package main

import "fmt"

type Dog struct{}

func (Dog) Speak() { fmt.Println("bowwow") }
func (Dog) Run()   { fmt.Println("yay,I'm running") }

type Cat struct{}

func (Cat) Speak() { fmt.Println("mewmew") }
func (Cat) Run()   { fmt.Println("shut up ****er") }

type Hoge struct{}

func (Hoge) Run() { fmt.Println("Hohho Hohho") }

// 処理に必要なものだけを宣言する
type Runner interface {
	Run()
}

func main() {
	Race([]Runner{Dog{}, Cat{}, Hoge{}})
}

func Race(runners []Runner) {
	for _, runner := range runners {
		runner.Run()
	}
}

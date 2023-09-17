package main

import (
	"errors"
	"fmt"
)

/*
// 前提として。以下がgolangのランタイムに組み込まれている
type error interface {
  Error() string
}
*/

// 埋め込む値のenumを作っておく
type Status int

const (
	HOGERA Status = iota
	FUGERI
	MOGERU
)

// エラーを保持できる構造体
type MyError struct {
	status  Status
	message string
}

// error interfaceにマッチさせるための関数
func (me MyError) Error() string {
	return me.message
}
func (me MyError) getStatus() Status {
	return me.status
}

// エラーを出す関数
func test() (string, MyError) {
	return "", MyError{
		status: HOGERA,
		//status:  FUGERI,
		message: "ふ、げらら えらら",
	}
}

func main() {
	// XXX どうなんだろ? なんか「割と常にswitch」でもいいんじゃなかろうか? と少し思えてきている
	switch _, err := test(); {
	case errors.As(err, &MyError{}):
		switch err.status {
		case HOGERA:
			fmt.Println("HOGERA")
			fmt.Println(err)
		case FUGERI:
			fmt.Println("FUGERI")
			fmt.Println(err)
		case MOGERU:
			fmt.Println("MOGERU")
			fmt.Println(err)
		}
	default:
		fmt.Println("エラーなんてなかったよ?")
	}
}

package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// いったん記号なし
var base = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 乱数のシードを設定する
func init() {
	// 重いけどセキュアだし強度が強い。ただ重すぎるなら time.Now().UnixNano() をseedにする、でもよいのかも？
	var seed int64
	if err := binary.Read(cryptorand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(err)
	}
	rand.Seed(seed)
}

// 指定した文字数のランダムな文字列を作成する
func MakeRandString1(num int) string {
	// スライス作る
	b := make([]byte, 0, num)

	// 乱数でぶんまわす
	l := len(base)
	for i := 0; i < num; i++ {
		b = append(b, base[rand.Intn(l)])
	}

	// 文字列にして返す
	return string(b)
}

func main() {
	s := MakeRandString1(20)
	fmt.Println(len(s), s)
}

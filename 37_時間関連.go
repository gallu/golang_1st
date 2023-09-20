package main

import (
	"fmt"
	"time"
)

func main() {
	// XXX こーゆー小技があるのか
	p := fmt.Println

	// 時間基本
	{
		// 色々あるぽ
		t := time.Now()
		p(t)
		p(t.Unix()) // これでエポック秒
		p(t.Year())
		p(t.Month())
		p(t.Day())
		p(t.Hour())
		p(t.Minute())
		p(t.Second())
		p(t.Nanosecond())
		p(t.Location())
		p(t.Weekday())

		// フォーマットの指定は、こう?
		p(t.Format(time.RFC3339)) // 2023-09-20T21:53:09+09:00
		// これでもいい(んだけど、値が「2006年1月2日15時4分5秒 アメリカ山地標準時MST(GMT-0700)」固定らしい)
		p(t.Format("2006-01-02 15:04:05Z07:00")) // 2023-09-20T21:53:09+09:00
	}

	// エポック
	{
		t := time.Unix(123456, 0)
		p(t)
		p(t.Unix()) // これでエポック秒
		p(t.Format(time.RFC3339))
	}
}

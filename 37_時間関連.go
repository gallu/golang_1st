package main

import (
	"fmt"
	"strconv"
	"strings"
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

		h, m, s := t.Clock()
		p("Clock", h, m, s)
	}

	// エポック
	{
		t := time.Unix(123456, 0)
		p(t)
		p(t.Unix()) // これでエポック秒
		p(t.Format(time.RFC3339))

		// ちょいと確認
		h, m, s := t.Clock()
		p("Clock", h, m, s)
		e1 := int64((h * 60 * 60) + (m * 60) + s)
		e2 := (t.Unix() + 9*60*60) % (24 * 60 * 60) // タイムゾーンあるので無理くり
		p(e1, e2, e1 == e2)

	}

	// シンプルに現在時刻
	{
		t := time.Now()
		fmt.Println("シンプルに現在時刻")
		fmt.Println(t.String())
		fmt.Println(t.Format(time.RFC3339))
	}
	// 指定日時(文字列から日付への変換)
	{
		t, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00+09:00")
		if err != nil {
			panic(err)
		}
		fmt.Println("指定日時(文字列から日付への変換)")
		fmt.Println(t.Format(time.RFC3339))
	}
	// 日付だけ(時間を切り捨て)
	{
		t := time.Now()
		fmt.Println("日付だけ(時間を切り捨て)")
		fmt.Println(t.Truncate(24 * time.Hour))
	}
	// 「日付が同じか」の比較
	{
		{
			t1 := time.Now()
			t2 := time.Now()
			fmt.Println("「日付が同じか」の比較")
			// fmt.Println(t1.Truncate(24 * time.Hour).Equal(t2.Truncate(24 * time.Hour))) // これはダメ: UTCベースで比較してる
			ts1 := t1.Format("2006-01-02")
			ts2 := t2.Format("2006-01-02")
			fmt.Println(ts1, ts2, ts1 == ts2)
		}
		{
			t1, _ := time.Parse(time.RFC3339, "2023-01-01T09:00:00+09:00")
			t2, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00+09:00")
			ts1 := t1.Format("2006-01-02")
			ts2 := t2.Format("2006-01-02")
			fmt.Println(ts1, ts2, ts1 == ts2)
		}
		{
			t1, _ := time.Parse(time.RFC3339, "2023-01-01T09:00:00+09:00")
			t2, _ := time.Parse(time.RFC3339, "2023-01-01T08:59:59+09:00")
			ts1 := t1.Format("2006-01-02")
			ts2 := t2.Format("2006-01-02")
			fmt.Println(ts1, ts2, ts1 == ts2)
		}
		{
			t1, _ := time.Parse(time.RFC3339, "2023-01-01T23:59:59+09:00")
			t2, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00+09:00")
			ts1 := t1.Format("2006-01-02")
			ts2 := t2.Format("2006-01-02")
			fmt.Println(ts1, ts2, ts1 == ts2)
		}
		{
			t1, _ := time.Parse(time.RFC3339, "2023-01-02T00:00:00+09:00")
			t2, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00+09:00")
			ts1 := t1.Format("2006-01-02")
			ts2 := t2.Format("2006-01-02")
			fmt.Println(ts1, ts2, ts1 == ts2)
		}
	}

	// UNIX秒への変換
	{
		t := time.Now()
		fmt.Println("UNIX秒への変換", t.Unix())
	}
	// UNIX秒からの変換
	{
		t := time.Unix(1700447200, 0)
		fmt.Println("UNIX秒からの変換", t.Format(time.RFC3339))
	}

	// (時間を無視して)日付の同一性確認
	{
		fmt.Println("(時間を無視して)日付の同一性確認")
		now, err := time.Parse(time.RFC3339, "2023-09-20T01:53:09+09:00")
		if err != nil {
			panic(err)
		}
		tests := []struct {
			ts string
			b  bool
		}{
			{
				ts: "2023-09-20T00:00:00+09:00",
				b:  true,
			},
			{
				ts: "2023-09-20T12:00:00+09:00",
				b:  true,
			},
			{
				ts: "2023-09-20T23:59:59+09:00",
				b:  true,
			},
			{
				ts: "2023-09-19T00:00:00+09:00",
				b:  false,
			},
			{
				ts: "2023-09-19T23:59:59+09:00",
				b:  false,
			},
			{
				ts: "2023-09-21T00:00:00+09:00",
				b:  false,
			},
		}
		for _, v := range tests {
			t, err := time.Parse(time.RFC3339, v.ts)
			if err != nil {
				panic(err)
			}
			// 「日付だけ」で比較
			t2 := t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
			nowTruc := now.Truncate(time.Hour).Add(-time.Duration(now.Hour()) * time.Hour)
			r := nowTruc.Equal(t2)
			fmt.Println(now.Format(time.RFC3339), ", ", t.Format(time.RFC3339))
			fmt.Println(nowTruc.Format(time.RFC3339), ", ", t2.Format(time.RFC3339), r, v.b)
		}
	}

	// 「指定された時刻」でリセットがかかるような諸々
	{
		fmt.Println("「指定された時刻」でリセットがかかるような諸々")
		// 「リセット時間」は与えられた情報
		resetTime := "04:05:06" // 夜中の4時5分6秒でリセットされる
		resetSec := clockToSec(resetTime)
		fmt.Println(resetTime, resetSec)

		//
		tests := []struct {
			now      string
			lastDraw string
			wont     bool
		}{
			// 「前日の指定時刻前」なら当然true
			{
				now:      "2023-09-20T05:53:09",
				lastDraw: "2023-09-19T01:53:09",
				wont:     true,
			},
			// 「前日の指定時刻後」でも、今が指定時刻後ならtrue
			{
				now:      "2023-09-20T05:53:09",
				lastDraw: "2023-09-19T01:53:09",
				wont:     true,
			},
			// 「当日の指定時刻前」でも、今が指定時刻後ならtrue
			{
				now:      "2023-09-20T05:53:09",
				lastDraw: "2023-09-20T01:53:09",
				wont:     true,
			},
			// 「当日の指定時刻後」で、今が指定時刻後ならfalse
			{
				now:      "2023-09-20T05:53:09",
				lastDraw: "2023-09-20T04:53:09",
				wont:     false,
			},
			// 境界線(指定時刻でリセットなので、現在時刻が1秒後ならOKになる)
			{
				now:      "2023-09-20T04:05:07",
				lastDraw: "2023-09-20T04:05:06",
				wont:     true,
			},
		}

		for _, v := range tests {
			// 現在時刻の取得
			// now := time.Now() // ほんとはこっち
			now, _ := time.Parse(time.RFC3339, v.now+"+09:00")
			// 今日の「日付だけ」
			nowDay := now.Truncate(time.Hour).Add(-time.Duration(now.Hour()) * time.Hour)
			// 今日の「時刻だけ」とその秒数
			nowClock := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
			nowSec := clockToSec(nowClock)

			// 確認
			fmt.Println(now.Format(time.RFC3339), nowDay.Format(time.RFC3339), nowClock, nowSec)

			/* 基準時刻(リセット年月日)の把握 */
			// いったん「現在日時」としておく
			pointTimeBase := nowDay
			if resetSec > nowSec {
				// 「今」がその時刻より未来なら「昨日のreset時刻」が基準点
				pointTimeBase = nowDay.Add(-24 * time.Hour) // 1日前
			}
			pointTime, _ := time.Parse(time.RFC3339, fmt.Sprintf("%sT%s+09:00", pointTimeBase.Format("2006-01-02"), resetTime))
			fmt.Println("pointTime:", pointTime.Format(time.RFC3339))

			// 判定ポイント(仮として「ログインがちゃを引いた時刻的な」)
			lastDraw, _ := time.Parse(time.RFC3339, v.lastDraw+"+09:00")

			// 比較
			r := pointTime.Compare(lastDraw) >= 0 // pointTime >= lastDraw
			fmt.Println("want:", v.wont, "get:", r)
			if v.wont != r {
				panic("不一致ぽ!!")
			}
		}

	}
}

// 「時刻の文字列」から「数値(秒)」に変換する
func clockToSec(s string) int {
	ss := strings.Split(s, ":")
	hour, _ := strconv.Atoi(ss[0])
	min, _ := strconv.Atoi(ss[1])
	sec, _ := strconv.Atoi(ss[2])
	all := (hour * 60 * 60) + (min * 60) + sec

	return all
}

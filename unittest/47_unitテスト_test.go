package unittest

import (
	"testing"
)

func Test最初のテストAdd(t *testing.T) {
	// 一番原始的な感じ?
	{
		got := Add(1, 2)
		want := 3
		if got != want {
			t.Errorf("get %d, want %d", got, want)
		}
	}
	// こんな書き方もあるっぽい(少しアレンジ)
	{
		// データの準備
		testData := []struct {
			name string
			args []int
			want int
		}{
			{
				name: "1, 2",
				args: []int{1, 2},
				want: 3,
			},
			{
				name: "111, 222",
				args: []int{111, 222},
				want: 333,
				//want: 334, // エラーにしてみる
			},
		}
		// ぶん回し
		for _, datum := range testData {
			got := Add(datum.args[0], datum.args[1])
			if got != datum.want {
				t.Errorf("%s: get %d, want %d", datum.name, got, datum.want)
			}
		}
	}
}

// publicじゃないメソッドでもテストできる？ できるぽい
func Test最初のテストhoge(t *testing.T) {
	// 一番原始的な感じ?
	{
		got := hoge(1)
		want := 2
		if got != want {
			t.Errorf("get %d, want %d", got, want)
		}
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// テスト用のURI(リクエストの内容を大体ざっくり出力してくれる phpコード)
	uri := "http://www.gjmj.net/http_test.php"

	// これが一番シンプルぽ?
	{
		resp, err := http.Get(uri)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	// getにパラメタ渡すならこんな感じ?
	{
		v := url.Values{}
		// SetとAddの2種類があるぽい
		v.Set("key_1", "value_1")
		v.Set("key_2", `data"'&=data`)

		resp, err := http.Get(uri + "?" + v.Encode())
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	// getに追加ヘッダ渡すならこんな感じ?
	{
		// resuest用意して
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("X-Hogera", "mugera fogera")

		// responseげと
		client := new(http.Client)
		resp, err := client.Do(req) // 多分ここでアクセスしてるんじゃなかろうか
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	// postを簡単に投げるならこれで
	{
		resp, err := http.Post(uri, "", nil)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	// データ添えてpost
	{
		v := url.Values{}
		// SetとAddの2種類があるぽい
		v.Set("key_1", "value_1")
		v.Set("key_2", `data"'&=data`)

		resp, err := http.PostForm(uri, v)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	// 追加ヘッダ付きでpost
	{
		v := url.Values{}
		v.Set("key_1", "value_1")
		v.Set("key_2", `data"'&=data`)

		// resuest用意して
		req, _ := http.NewRequest("POST", uri, strings.NewReader(v.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // これちゃんと設定しないとデータとどかないぽ?
		req.Header.Set("X-Hogera", "mugera fogera")

		// responseげと
		client := new(http.Client)
		resp, err := client.Do(req) // 多分ここでアクセスしてるんじゃなかろうか
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // なんか閉じるのか……通信してるからそれのcloseかしらん?

		//
		fmt.Println("Status", resp.Status, resp.StatusCode)

		// これで読めるぽ
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ミニマム構成、こんなもん?
	//  curl localhost:8080/hello で動いた
	{
		/*
			// 「入ってくるURL」と動作の指定
			hello := func(w http.ResponseWriter, req *http.Request) {
				fmt.Fprintf(w, "hello\n")
			}
			http.HandleFunc("/hello", hello)
			// サーバ起動
			http.ListenAndServe(":8080", nil)
		*/
	}

	// こーゆー設定もあるぽ
	{
		// 実際には色々とパラメタがあるぽ
		server := http.Server{
			Addr:    ":8080",
			Handler: nil,
		}

		// 「入ってくるURL」と動作の指定
		hello := func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "hello\n")
			fmt.Fprintln(w, "hoge", req.FormValue("hoge"))
		}
		http.HandleFunc("/hello", hello)
		// サーバ起動
		server.ListenAndServe()
	}

}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//
	uri := "https://www.m-fr.net"

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

}

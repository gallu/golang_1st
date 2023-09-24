package main

import (
	"encoding/json"
	"fmt"
)

type Parts struct {
	Id   int
	Name string
}

type Whole struct {
	Id   int
	Name string
	Arr  []int
	Data []Parts
}

func main() {
	// 超単純なjson
	{
		// json文字列
		json_string := `{
			"id": 123,
			"name": "gallu"
		}`
		// 格納用の構造体(ラフに用意)
		st := struct {
			Id   int
			Name string
		}{}
		// これで分解できる?
		err := json.Unmarshal([]byte(json_string), &st)
		if err != nil {
			panic(err)
		}

		fmt.Println(st)
		fmt.Println(st.Id, st.Name)
	}
	// 別解?
	{
		// json文字列
		json_string := `{
			"id": 123,
			"name": "gallu"
		}`
		// 格納用のmap(ラフに用意)
		//var dat map[string]interface{}
		var dat map[string]any

		// これで分解できる?
		err := json.Unmarshal([]byte(json_string), &dat)
		if err != nil {
			panic(err)
		}

		fmt.Println(dat)
		fmt.Println(dat["id"], dat["name"])
		fmt.Printf("%T, %T \n", dat["id"], dat["name"])
		id := int(dat["id"].(float64)) // 一応これでいけるけど、遠回しだねぇ……(直接intへの変換、は出来なかった……)
		fmt.Println(id)
	}
	// 配列入り
	{
		// 生json
		json_string := `{"id":123,"name":"gallu","arr":[1,2,3],"data":[{"id":1,"name":"\u30c0\u30ac\u30fc"},{"id":2,"name":"\u5263"},{"id":3,"name":"\u30c0\u30d6\u30eb\u30a2\u30c3\u30af\u30b9"}]}`
		// 格納用の構造体
		whole := Whole{}
		// 分解
		err := json.Unmarshal([]byte(json_string), &whole)
		if err != nil {
			panic(err)
		}

		fmt.Println(whole)
	}
	// ざっくりなら、これでいける?
	{
		// 生json
		json_string := `{"id":123,"name":"gallu","arr":[1,2,3],"data":[{"id":1,"name":"\u30c0\u30ac\u30fc"},{"id":2,"name":"\u5263"},{"id":3,"name":"\u30c0\u30d6\u30eb\u30a2\u30c3\u30af\u30b9"}]}`
		// 格納用のmap(ラフに用意)
		var dat map[string]any
		err := json.Unmarshal([]byte(json_string), &dat)
		if err != nil {
			panic(err)
		}
		// いけるぽ
		fmt.Println(dat)
	}
	// データが欠けてたら?
	// 普通に取れるぽい。だとすると「構造体には満額をパンパンに定義」しておくのがよろしのかしらん? 多分「ゼロ値」が入ってるぽいから
	{
		// 生json
		json_string := `{"id":123,"data":[{"id":1,"name":"\u30c0\u30ac\u30fc"},{"id":2,"name":"\u5263"},{"id":3,"name":"\u30c0\u30d6\u30eb\u30a2\u30c3\u30af\u30b9"}]}`
		// 格納用の構造体
		whole := Whole{}
		// 分解
		err := json.Unmarshal([]byte(json_string), &whole)
		if err != nil {
			panic(err)
		}

		fmt.Println(whole)
		fmt.Println(whole.Name)
	}

	// 書き出し側
	{
		parts := Parts{
			Id:   123,
			Name: "hoge",
		}
		// ここもbyteなんだ。なるほど。
		json_string, err := json.Marshal(parts)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json_string))
	}
}

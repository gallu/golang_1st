package main

import (
	"encoding/json"
	"fmt"
)

type AnyArray map[string]interface{}

func main() {
	//
	{
		d := make(AnyArray)
		d["key1"] = 123
		d["key2"] = "string"
		d["key3"] = true

		dd := make(AnyArray)
		dd["key4-1"] = 987
		dd["key4-2"] = "hoge"
		d["key4"] = dd
		// d["key4"]["key4-3"] = true //  (map index expression of type interface{}) でエラーになる
		d["key4"].(AnyArray)["key4-3"] = true // 型アサーションすればいける(けど、書式的にどうなん？？？)

		json_string, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", d)
		fmt.Println(string(json_string))
	}
}

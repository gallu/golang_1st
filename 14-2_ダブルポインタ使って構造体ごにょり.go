package main

import "fmt"

type Hoge struct {
	Id int
}

var List []*Hoge
var BaseList []**Hoge
var UseList []**Hoge

func main() {
	//
	{
		// まずデータ突っ込む
		for i := 1; i < 5; i++ {
			h := &Hoge{Id: i}
			List = append(List, h)
			BaseList = append(BaseList, &h)
			UseList = append(UseList, &h)
		}
		fmt.Printf("List: %+v \n", List)
		fmt.Printf("BaseList: %+v \n", BaseList)
		fmt.Printf("UseList: %+v \n", UseList)
		// 確認
		for i, v := range UseList {
			fmt.Printf("%d: %#v\n", i, (*v))
		}
		fmt.Println("")

		// 部分的に削除
		for _, v := range BaseList {
			if (*v).Id == 2 {
				(*v) = nil
			} else if (*v).Id == 3 {
				(*v) = nil
			}
		}
		for i, v := range UseList {
			if (*v) == nil {
				continue
			}
			fmt.Printf("%d: %#v\n", i, (*v))
		}
		fmt.Printf("List: %+v \n", List)
		fmt.Printf("BaseList: %+v \n", BaseList)
		fmt.Printf("UseList: %+v \n", UseList)
	}
}

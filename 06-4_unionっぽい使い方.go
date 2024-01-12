package main

import (
	"encoding/json"
	"fmt"
)

type IIDD struct {
	Id   *int
	Name *string
}

func NewIIDDFromId(i int) *IIDD {
	return &IIDD{
		Id: &i,
	}
}

func NewIIDDFromName(n string) *IIDD {
	return &IIDD{
		Name: &n,
	}
}

type Hoge struct {
	Id   *IIDD
	Data string
}

func NewHogeFromId(i int, d string) *Hoge {
	return &Hoge{
		Id:   NewIIDDFromId(i),
		Data: d,
	}
}
func NewHogeFromName(n string, d string) *Hoge {
	return &Hoge{
		Id:   NewIIDDFromName(n),
		Data: d,
	}
}

func main() {
	{
		h := NewHogeFromId(1, "hoge")
		s, err := json.Marshal(h)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s))
	}
	{
		h := NewHogeFromName("name", "foo")
		s, err := json.Marshal(h)
		if err != nil {
			panic(err)
		}
		fmt.Println("json: ", string(s))
	}
}

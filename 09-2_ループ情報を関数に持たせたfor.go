package main

import "fmt"

type Hoge struct {
	data map[string]string
	pos  map[string]string
}

func NewHoge(d map[string]string) *Hoge {
	return &Hoge{
		data: d,
	}
}
func (h *Hoge) Start() {
	h.pos = make(map[string]string)
	for k, v := range h.data {
		h.pos[k] = v
	}
}
func (h *Hoge) Get() (map[string]string, string, string) {
	r2 := h.pos
	// 雑に１要素を把握
	key := ""
	r1 := ""
	for k, _ := range h.pos {
		key = k
	}
	if key != "" {
		r1 = h.pos[key]
		delete(h.pos, key)
	} else {
		r1 = ""
	}
	return r2, key, r1
}

func main() {
	// いけるけど、あんまり好みじゃないなぁ……
	{
		h := NewHoge(
			map[string]string{
				"key":  "10",
				"hoge": "999",
				"foo":  "1234",
			},
		)
		fmt.Printf("%+v \n", h)
		h.Start()
		fmt.Printf("%+v \n", h)
		for {
			pos, key, val := h.Get()
			fmt.Printf("%+v, %+v, %+v \n", key, val, pos)
			if len(pos) == 0 {
				break
			}
		}
	}
	println("")
	// こっちのが「まし」だなぁ
	{
		f := NewFoo(
			map[string]string{
				"key":  "10",
				"hoge": "999",
				"foo":  "1234",
			},
		)
		for i := 0; i < 5; i++ {
			f.Start()
			for f.Next() != false {
				pos, key, val := f.Get()
				fmt.Printf("%+v, %+v, %+v \n", key, val, pos)
			}

		}
	}

}

type Foo struct {
	data map[string]string
	pos  map[string]string
}

func NewFoo(d map[string]string) *Foo {
	return &Foo{
		data: d,
	}
}
func (f *Foo) Start() {
	f.pos = make(map[string]string)
	for k, v := range f.data {
		f.pos[k] = v
	}
}
func (f *Foo) Next() bool {
	return f.pos != nil
}

func (f *Foo) Get() (map[string]string, string, string) {
	r2 := f.pos
	// 雑に１要素を把握
	key := ""
	r1 := ""
	for k, _ := range f.pos {
		key = k
	}
	if key != "" {
		r1 = f.pos[key]
		delete(f.pos, key)
		if len(f.pos) == 0 {
			f.pos = nil
		}
	} else {
		r1 = ""
	}
	return r2, key, r1
}

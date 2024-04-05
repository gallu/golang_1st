package main

import "fmt"

type Hoge struct {
	id int
}

func NewHoge(i int) **Hoge {
	r := &Hoge{
		id: i,
	}
	return &r
}

func main() {
	main := []**Hoge{}
	sub1 := []**Hoge{}

	h1 := NewHoge(1)
	h2 := NewHoge(2)

	main = append(main, h1)
	main = append(main, h2)

	sub1 = append(sub1, h1)
	sub1 = append(sub1, h2)

	fmt.Printf("%+v \n", *main[0])
	fmt.Printf("%+v \n", *sub1[0])

	h := main[0]
	// これ、どっちがいいんだろうねぇ？
	fmt.Println((*h).id)
	fmt.Println((**h).id)

	// 消去
	// *main[0] = nil

	// h := main[0]
	*h = nil

	fmt.Printf("%+v \n", *main[0])
	fmt.Printf("%+v \n", *sub1[0])
}

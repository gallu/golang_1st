package main

import "fmt"

type Hoge struct {
	Id   int
	Name string
}

type HogeI interface {
	GetName() string
	SetName(s string)
}

func NewHoge(id int, name string) Hoge {
	return Hoge{
		Id:   id,
		Name: name,
	}
}

func (h *Hoge) GetName() string {
	return h.Name
}
func (h *Hoge) SetName(s string) {
	h.Name = s
}

type Foo struct {
	Id int
	HogeI
}

type Bar struct {
	Hoge
}

func main() {
	{
		f := Foo{
			Id: 1,
			HogeI: &Hoge{ // インタフェースだと、そもそもポインタにしない限りはいらないのか(&削ったらエラーになった)
				Id:   999,
				Name: "hoge",
			},
		}
		fmt.Printf("%+v \n", f)
	}

	{
		h := NewHoge(1, "hoge name")
		f := Foo{
			Id:    1,
			HogeI: &h, // ここに NewHoge() 直、はできない
		}
		fmt.Printf("%+v \n", f)
	}

	{
		f := Foo{
			Id: 1,
			HogeI: &Hoge{ // インタフェースだと、そもそもポインタにしない限りはいらないのか(&削ったらエラーになった)
				Id:   999,
				Name: "hoge",
			},
		}
		// 当然ながら「ポインタの値は違う」んだけど、どうも「内在しているものは同じ」っぽい。から、f.HogeIは同じところを見ている
		fmt.Printf("main\t%+v, %p, %p\n", f, &f, f.HogeI)
		p1(f)
		// ポインタで渡ってるから関数先での変更がこっちに影響している
		fmt.Printf("main\t%+v\n", f.HogeI)
	}

	{
		b := Bar{
			Hoge: Hoge{
				Id:   1,
				Name: "bar hoge",
			},
		}
		// このパターンだと、多分、copyされてる
		fmt.Printf("main\t%+v, %p, %p\n", b, &b, &b.Hoge)
		p2(b)
		// 実態がcopyされているから関数先の変更とかしらんぷいぷい
		fmt.Printf("main\t%+v\n", b.Hoge)
	}
}

func p1(f Foo) {
	fmt.Printf("p1\t%+v, %p, %p\n", f, &f, f.HogeI)
	f.SetName("p1 set")
}

func p2(b Bar) {
	fmt.Printf("p2\t%+v, %p, %p\n", b, &b, &b.Hoge)
	b.Hoge.Name = "p2 set"
}

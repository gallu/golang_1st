package main

import "fmt"

type Percentage uint

func NewPercentage(v float64) Percentage {
	v += 0.00000001
	fmt.Println(v)
	return Percentage(v * 100)
}
func (p Percentage) Value() float64 {
	return float64(p) / 100
}

type Hoge struct {
	Str Percentage
}

func (h *Hoge) GetStr() float64 {
	return float64(h.Str) / 100
}

func NewHoge(str float64) *Hoge {
	return &Hoge{
		Str: NewPercentage(str),
	}
}

func main() {
	{
		f1 := 2.2
		f2 := 18.06
		// h := NewHoge(2.2 + 18.06)
		h := NewHoge(f1 + f2)
		fmt.Printf("%+v\n", h)
		fmt.Printf("%.32f \n", 2.2+18.06)
		fmt.Printf("%.32f \n", 20.26)
		fmt.Printf("%.32f \n", f1)
		fmt.Printf("%.32f \n", f2)
		fmt.Printf("%.32f \n", f1+f2)
		fmt.Printf("\n")
		fmt.Printf("Str: %+v\n", h.GetStr())
	}
	{
		i := Percentage(10)
		j := Percentage(20)
		k := i + j
		fmt.Printf("%T, %v \n", k, k)
	}
}

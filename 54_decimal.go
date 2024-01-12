package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// n := decimal.NewFromFloat(123.44444)
	// n2 := decimal.NewFromFloat(123.44444)
	n := decimal.NewFromFloat(123.111111)
	n2 := decimal.NewFromFloat(123.1111111)
	n3 := n.Add(n2)
	// r := n3.RoundBank(2) // 四捨五入
	// r := n3.RoundDown(2) // 切り捨て
	r := n3.RoundUp(2) // 切り上げ

	fmt.Println(n.String(), n2.String(), n3.String(), r.String()) //

	f, exact := r.Float64()
	fmt.Printf("%T, %v: %v \n", f, f, exact)
}

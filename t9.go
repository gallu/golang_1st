package main

import "fmt"

func main() {
	{
		f1 := 2.2
		f2 := 18.06
		fmt.Printf("%v, %v, %v \n", f1, f2, f1+f2)
		fmt.Printf("%.32f\n", f1)
		fmt.Printf("%.32f\n", f2)
		fmt.Printf("%.32f\n", f1+f2)
	}

	{
		f1 := 2.2 * 100
		f2 := 18.06 * 100
		fmt.Printf("%v, %v, %v \n", f1/100, f2/100, (f1+f2)/100)
	}
}

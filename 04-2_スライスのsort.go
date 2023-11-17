package main

import (
	"fmt"
	"sort"
)

func main() {
	// 普通っぽくsort
	{
		arr := []int{5, 1, 13, 4, 6, 8, 5, 9, 7, 3}
		fmt.Println(arr)

		// 昇順
		sort.Slice(arr, func(i, j int) bool {
			fmt.Println("i:", i, arr[i], "j:", j, arr[j])
			return arr[i] < arr[j]
		})
		fmt.Println(arr)

		// 降順(安定ソート)
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		fmt.Println(arr)

		// 変な条件で
		sort.Slice(arr, func(i, j int) bool {
			if arr[i] == 5 {
				return true
			}
			if arr[i] == 7 {
				return false
			}
			if arr[i] == 9 {
				return true
			}
			return arr[i] > arr[j]
		})
		fmt.Println(arr)
	}
}

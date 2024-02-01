package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr[:])))
	
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

}
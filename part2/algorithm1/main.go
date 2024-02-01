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

	sort.Ints(arr[:])
	
	fmt.Println(arr[9])
	fmt.Println(arr[8])
	fmt.Println(arr[7])

}
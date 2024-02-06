package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)

	arr := make([]int, n)

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ") 
	for i := 0; i < len(inputs); i++ {
		input, _ := strconv.Atoi(inputs[i])
		arr[i] = input
	}

	count := 0

	for i := 0; i < n - 1; i++ {
		min_i := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_i] {
				min_i = j
			}
		}
		arr[i], arr[min_i] = arr[min_i], arr[i]
		if i != min_i {
			count++
		}
	}

	fmt.Println(arr)
	fmt.Println(count)
}
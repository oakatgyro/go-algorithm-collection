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
	
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
		fmt.Println(arr)
	}
}
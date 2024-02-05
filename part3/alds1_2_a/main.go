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

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count++
			}
		}
	}
	
	arrString := make([]string, n)
	for i := 0; i < n; i++ {
		arrString[i] = strconv.Itoa(arr[i])
	}
	fmt.Println(strings.Join(arrString, " "))
	fmt.Println(count)
}
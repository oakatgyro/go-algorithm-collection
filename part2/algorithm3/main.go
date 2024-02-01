package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var arr []int
	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		arr = append(arr, p)
	}

	for i := len(arr); i > len(arr)-3; i-- {
		fmt.Println(arr[i-1])
	}

}
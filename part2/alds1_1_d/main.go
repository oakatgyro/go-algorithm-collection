package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)

	min := math.MaxInt
	max := -math.MaxInt

	for i := 0; i < n; i++ {
		sc.Scan()
		input, _ := strconv.Atoi(sc.Text())

		if input < min {
			min = input
		}

		if input > max {
			max = input
		}
	}
	
	fmt.Println(max-min)
}
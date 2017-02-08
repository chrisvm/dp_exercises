package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		sum = 14
		n   = 3
	)

	var (
		coins = [n]int{1, 3, 5}
		min   = [sum + 1]int{}
	)

	// set min[i] = inf for all i
	min[0] = 0
	for index := 1; index <= sum; index++ {
		min[index] = math.MaxInt32
	}

	for i := 1; i <= sum; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i && min[i-coins[j]]+1 < min[i] {
				min[i] = min[i-coins[j]] + 1
			}
		}
	}

	fmt.Printf("coins = %d\n", min[sum])
}

package main

import (
	"fmt"
)

func plusMinus(arr []int) {
	n := len(arr)
	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for _, value := range arr {
		if value > 0 {
			positiveCount++
		} else if value < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	positiveRatio := float64(positiveCount) / float64(n)
	negativeRatio := float64(negativeCount) / float64(n)
	zeroRatio := float64(zeroCount) / float64(n)

	fmt.Printf("%.6f\n", positiveRatio)
	fmt.Printf("%.6f\n", negativeRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}

func main() {
	arr := []int{1, 1, 0, -1, -1}
	plusMinus(arr)
}

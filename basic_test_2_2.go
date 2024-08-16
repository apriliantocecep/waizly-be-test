package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter number of elements:")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter element:")
		fmt.Scan(&arr[i])
	}

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

package main

import (
	"fmt"
	"sort"
)

// Problem Solving Basic - Test 1

func miniMaxSum(arr []int) {
	sort.Ints(arr) // Mengurutkan array
	minSum := 0
	maxSum := 0

	// Menghitung jumlah minimum dengan menjumlahkan 4 elemen pertama
	for i := 0; i < 4; i++ {
		minSum += arr[i]
	}

	// Menghitung jumlah maksimum dengan menjumlahkan 4 elemen terakhir
	for i := 1; i < 5; i++ {
		maxSum += arr[i]
	}

	fmt.Println(minSum, maxSum)
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	miniMaxSum(arr)
}

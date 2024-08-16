package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input lima angka yang dipisahkan oleh spasi
	var arr [5]int64
	//for i := 0; i < 5; i++ {
	//	fmt.Scan(&arr[i])
	//}

	// Input lima angka yang dipisahkan oleh spasi
	fmt.Scanf("%d %d %d %d %d", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4])

	// Mengurutkan array
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Menghitung jumlah minimum dan maksimum
	var minSum, maxSum int64
	for i := 0; i < 4; i++ {
		minSum += arr[i]
	}
	for i := 1; i < 5; i++ {
		maxSum += arr[i]
	}

	// Menampilkan hasil
	fmt.Println(minSum, maxSum)
}

package main

import (
	"fmt"
	"math/rand"
)

func selectionSort(arr []int) {
	// Sort an arr using Selection Sort
	fmt.Println("Original:", arr)

	// Loop through the array starting from the smallest index
	for i := 0; i < len(arr); i++ {
		// Find the min value and swap it with i
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		// Swap min and i
		arr[i], arr[min] = arr[min], arr[i]
		fmt.Printf("%v: %v\n", i, arr)

	}
}

func insertionSort(arr []int) {
	// Sort an arr using Insertion Sort
	fmt.Println("Original:", arr)

	N := len(arr)
	compares := 0

	// Loop through the array
	for i := 1; i < N; i++ {
		// Swap consecutive inversions
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			compares++
		}
		fmt.Printf("%v: %v\n", i, arr)
	}
	fmt.Println("Total number of compares: ", compares)

}

func shellSort(arr []int) {
	// Sort an arr using Shell Sort
	fmt.Println("Original", arr)

	N := len(arr)
	h := 1
	compares := 0

	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40
	}

	for h >= 1 {
		// h-sort the array
		for i := h; i < N; i++ {
			// Insert a[i] among a[i-h], a[i-2*h], a[i-3*h]...
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j-h], arr[j] = arr[j], arr[j-h]
				compares++
			}
			fmt.Printf("%v %v: %v\n", h, i, arr)
		}
		h = h / 3
		fmt.Printf("%v: %v\n", h, arr)
	}

	fmt.Println("Total number of compares: ", compares)
}

func main() {
	// Create a array from 1 to N
	N := 16
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i + 1
	}

	// Shuffle the values
	rand.Shuffle(N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	// Selection Sort
	cp := make([]int, N)
	copy(cp, arr)
	fmt.Println("Selection Sort")
	selectionSort(cp)

	// Insertion Sort
	copy(cp, arr)
	fmt.Println("Insertion Sort")
	insertionSort(cp)

	// Shell Sort
	copy(cp, arr)
	fmt.Println("Shell Sort")
	shellSort(cp)
}

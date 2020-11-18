package main

import (
	"fmt"
	"math/rand"
)

func absInplaceSort(arr, aux []int, lo, mid, hi int) {
	// Merge a[lo...mid] with a[mid+1...hi]
	i, j := lo, mid+1

	copy(aux, arr)

	for k := lo; k <= hi; k++ {
		// Merge back to a[lo...hi]
		switch {
		case i > mid:
			// Exchausted options for i
			arr[k] = aux[j]
			j++
		case j > hi:
			// Exchausted options for j
			arr[k] = aux[i]
			i++
		case aux[j] < aux[i]:
			// j is less than i, get j
			arr[k] = aux[j]
			j++
		default:
			// i is less than j, get i
			arr[k] = aux[i]
			i++
		}
	}
}

func topDownSort(a, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	fmt.Println("Sorting left: ", a[lo:mid+1])
	fmt.Println("Sorting right: ", a[mid+1:hi+1])
	topDownSort(a, aux, lo, mid)        // Sort left half
	topDownSort(a, aux, mid+1, hi)      // Sort right half
	absInplaceSort(a, aux, lo, mid, hi) // Merge results
	fmt.Printf("Merged lo: %v and hi: %v %v\n", lo, hi, a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bottomUpSort(a, aux []int) {
	// Do lg N passes of pairwise merges
	N := len(a)

	// Loop through subarray sizes
	for sz := 1; sz < N; sz += sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			absInplaceSort(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1))
		}
		fmt.Printf("Sub Array Size %v: %v\n", sz, a)
	}

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

	// Abs in place test
	test := []int{4, 5, 3, 14}
	auxTest := make([]int, len(test))
	copy(auxTest, test)
	lo, hi := 0, len(test)-1
	mid := lo + (hi-lo)/2
	absInplaceSort(test, auxTest, lo, mid, hi)
	fmt.Println("Abs In Place Test")
	fmt.Println("Original: ", auxTest)
	fmt.Println("Sorted: ", test)

	// Top-down sort
	cp := make([]int, N)
	copy(cp, arr)
	fmt.Printf("\nTop-Down\n")
	fmt.Println("Original: ", cp)
	topDownSort(cp, arr, 0, N-1)
	fmt.Println("Sorted: ", cp)

	// Bottom-up sort
	copy(cp, arr)
	fmt.Printf("\nBottom-Up\n")
	fmt.Println("Original: ", cp)
	bottomUpSort(cp, arr)
	fmt.Println("Sorted: ", cp)
}

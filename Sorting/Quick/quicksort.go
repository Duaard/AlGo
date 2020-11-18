package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int, lo, hi int) {
	if hi <= lo {
		return
	}

	j := partition(a, lo, hi)
	quickSort(a, lo, j-1)
	quickSort(a, j+1, hi)
}

func partition(a []int, lo, hi int) int {
	// Partition into a[lo..i-1], a[i], a[i+1..hi]
	i, j := lo+1, hi // left and right scan indices
	v := a[lo]       // partitioning item

	for {
		// Scan right, scan left, check for scan complete then exchange
		for ; a[i] < v; i++ {
			if i == hi {
				break
			}
		}
		for ; v < a[j]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo] // Put v = a[j] into position
	fmt.Printf("Partition %v: %v\n", v, a[lo:hi+1])
	return j // with a[lo..j-1] <= a[j] <= a[j+1..hi]
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

	// Quick Sort
	cp := make([]int, N)
	copy(cp, arr)
	lo, hi := 0, len(cp)-1
	fmt.Println("Original: ", cp)
	quickSort(cp, lo, hi)
	fmt.Println("Sorted: ", cp)

}

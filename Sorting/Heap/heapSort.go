package main

// Modified sink function used in heap sort
func heapSink(a []int, i, n int) {
	for 2*i <= n {
		c := 2 * i // Get children

		// fmt.Printf("c: %v n: %v\n", c, n)
		if c < n && a[c] < a[c+1] {
			c++
		}
		if a[i] >= a[c] {
			break
		}

		// Exchange with child
		a[i], a[c] = a[c], a[i]
		i = c
	}
}

// Sort given array using sink
func heapSort(a []int) {
	n := len(a) - 1

	for k := n / 2; k >= 1; k-- {
		heapSink(a, k, n)
	}

	for n > 1 {
		a[n], a[1] = a[1], a[n]
		n--
		heapSink(a, 1, n)
	}
}

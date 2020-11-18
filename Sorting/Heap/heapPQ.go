package main

// A priority queue implementation using a heap tree
type HeapPriorityQueue struct {
	N  int   // Keeps track of current size
	pq []int // Heap-ordered complete binary tree
}

// Reheapify using bottom up approach
func (h *HeapPriorityQueue) swim(i int) {
	for i > 1 && h.pq[i] > h.pq[i/2] {
		// Exchange with parent
		h.pq[i/2], h.pq[i] = h.pq[i], h.pq[i/2]
		i = i / 2
	}

}

// Reheapify using top down approach
func (h *HeapPriorityQueue) sink(i int) {
	for 2*i <= h.N {
		c := 2 * i // Get children

		if c < h.N && h.pq[c] < h.pq[c+1] {
			c++
		}
		if h.pq[i] >= h.pq[c] {
			break
		}

		// Exchange with child
		h.pq[i], h.pq[c] = h.pq[c], h.pq[i]
		i = c
	}
}

type emptyQueue struct{}

func (err *emptyQueue) Error() string {
	return "Priority queue is empty"
}

// Returns the size of the queue
func (h *HeapPriorityQueue) size() int {
	return h.N
}

// Returns max value in queue
func (h *HeapPriorityQueue) top() (int, error) {
	if h.size() > 0 {
		return h.pq[1], nil
	}
	return 0, &emptyQueue{}
}

// Adds given int to queue and then reheapify
func (h *HeapPriorityQueue) push(v int) {
	// Check if priority queue is empty,
	// if so add a filler to skip 0 index
	if h.pq == nil {
		h.pq = append(h.pq, 0)
	}
	h.pq = append(h.pq, v)
	h.N++
	h.swim(h.N)
}

// Removes the largest key and then reheapify
func (h *HeapPriorityQueue) pop() {
	// Check if priority queue is empty,
	// if so just return
	if h.size() == 0 {
		return
	}

	// Exchange root with the last item
	h.pq[1] = h.pq[h.N]
	// Resize the slice
	h.pq = h.pq[:h.N]
	h.N--

	h.sink(1)
}

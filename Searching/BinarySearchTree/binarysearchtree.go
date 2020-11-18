package main

import "fmt"

type Node struct {
	key   Key   // Key used for node
	value Value // Associated value
	left  *Node // Link to left node
	right *Node // Link to right node
	count int   // Num of nodes in subtree
}

// BinarySearchTree implementation of a Symbol Table
type BinarySearchTree struct {
	root    *Node
	compare func(interface{}, interface{}) int // Allow client to define a valid comparitor
}

// Helper function to get size of node
func size(n *Node) int {
	if n == nil {
		return 0
	}
	left := size(n.left)
	right := size(n.right)
	return left + right + 1
}

// Get the value asociated with key
func (b *BinarySearchTree) Get(key Key) Value {
	return get(b.root, key, b.compare)
}

func get(n *Node, key Key, compare func(interface{}, interface{}) int) Value {
	// Reached end of recursion, key doesn't exist
	if n == nil {
		return nil
	}

	// Use comparator to look for the key recursively
	switch compare(key, n.key) {
	case -1:
		// The key is in the left tree
		return get(n.left, key, compare)
	case 1:
		// The key is in the right tree
		return get(n.right, key, compare)
	default:
		// Key found return appropriate value
		return n.value
	}
}

// Put the value associated with key
func (b *BinarySearchTree) Put(key Key, val Value) {
	// Search for key, update value if found and grow if new
	b.root = put(b.root, key, val, b.compare)
}

func put(n *Node, key Key, val Value, f func(interface{}, interface{}) int) *Node {
	// Reached the end of traversal, add in new key
	if n == nil {
		return &Node{key: key, value: val, count: 1}
	}

	// Use comparator to look for the key recursively
	switch f(key, n.key) {
	case -1:
		// Key is less than current key
		n.left = put(n.left, key, val, f)
	case 1:
		// Key is greater than current key
		n.right = put(n.right, key, val, f)
	default:
		// Key is found, set approp value
		n.value = val
	}
	// Update this n's count
	n.count = size(n.left) + size(n.right) + 1
	return n
}

// Print all contents of the binary search tree
func (b *BinarySearchTree) Print() {
	print(b.root)
}

func print(n *Node) {
	// Print using DFS
	if n.left != nil {
		print(n.left)
	}
	fmt.Printf("Key: %v Val: %v Count: %v\n", n.key, n.value, n.count)
	if n.right != nil {
		print(n.right)
	}
}

// Min returns the min key in the tree
func (b *BinarySearchTree) Min() Key {
	return min(b.root).key
}

// Get the left most value
func min(n *Node) *Node {
	// No more left value, return n
	if n.left == nil {
		return n
	}

	// Get min of left
	return min(n.left)
}

// Max returns the max key in the tree
func (b *BinarySearchTree) Max() Key {
	return max(b.root).key
}

// Get the right most value
func max(n *Node) *Node {
	// No more right value, return n
	if n.right == nil {
		return n
	}

	// Get max of right
	return max(n.right)
}

// Floor returns the largest Key less than given Key
func (b *BinarySearchTree) Floor(k Key) Key {
	return floor(b.root, k, b.compare).key
}

// Helper function to allowe recursion
func floor(n *Node, k Key, f func(interface{}, interface{}) int) *Node {
	if n == nil {
		return nil
	}
	var t *Node
	switch f(k, n.key) {
	case -1:
		// The floor is in the left
		return floor(n.left, k, f)
	case 1:
		// The floor could be in the right
		t = floor(n.right, k, f)
	default:
		// Return n
		return n
	}

	if t != nil {
		return t
	}

	return n
}

// Ceil returns the smallest Key larger than give Key
func (b *BinarySearchTree) Ceil(k Key) Key {
	return ceil(b.root, k, b.compare).key
}

// Helper function to allow recursion
func ceil(n *Node, k Key, f func(interface{}, interface{}) int) *Node {
	if n == nil {
		return nil
	}

	var t *Node

	switch f(k, n.key) {
	case -1:
		// The key might be in the left tree
		t = ceil(n.left, k, f)
	case 1:
		// The key is in the right tree
		return ceil(n.right, k, f)
	default:
		// Found an equal key
		return n
	}

	if t != nil {
		return t
	}
	return n
}

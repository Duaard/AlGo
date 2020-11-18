package main

import (
	"fmt"
	"testing"
)

// Test binarysearchtree Put function
func TestPut(t *testing.T) {
	f := func(a interface{}, b interface{}) int {
		x, y := a.(int), b.(int)
		switch {
		case x > y:
			return 1
		case x < y:
			return -1
		default:
			return 0
		}
	}

	bst := &BinarySearchTree{compare: f}

	bst.Put(5, 10)
	bst.Put(7, 3)
	bst.Put(2, 6)
	bst.Put(5, 15)
	bst.Put(10, 20)
	bst.Put(15, 25)
	bst.Put(1, 3)
	bst.Put(3, 69)
	bst.Print()

	fmt.Println("Min key: ", bst.Min())
	fmt.Println("Max key: ", bst.Max())
	val := 4
	fmt.Printf("Floor of %v: %v\n", val, bst.Floor(val))
	fmt.Printf("Ceil of %v: %v\n", val, bst.Ceil(val))
}

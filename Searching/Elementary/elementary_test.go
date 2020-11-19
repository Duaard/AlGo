package main

import (
	"fmt"
	"testing"
)

func f(a interface{}, b interface{}) int {
	x, y := a.(int), b.(int)
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	default:
		return 0
	}

}

func TestElementary(t *testing.T) {
	// Create a new SymbolTable
	st := &SymbolTable{compare: f}
	st.New()

	st.Put(5, 10)
	st.Put(10, 5)
	fmt.Println(st.Get(5))

	st.Put(5, 15)
	st.Put(1, 3)
	st.Print()
}

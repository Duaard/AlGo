package main

import "reflect"

// CompareGeneric compares 2 empty interfaces, caller must provide proper comparator function
func CompareGeneric(a interface{}, b interface{}, f func(interface{}, interface{}) int) int {
	t1, t2 := reflect.TypeOf(a), reflect.TypeOf(b)

	if t1 != t2 {
		// Error, comparing different types
		panic("Comparing different types.")
	}

	return f(a, b)

}

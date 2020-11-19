package main

import "fmt"

type Key interface {
}

type Value interface{}

type SymbolTable struct {
	keys    []Key
	values  []Value
	compare func(interface{}, interface{}) int // Client provided way to compare Keys
}

func (s *SymbolTable) New() {
	// Initialize keys and values
	s.keys = make([]Key, 0)
	s.values = make([]Value, 0)
}

func (s *SymbolTable) Put(k Key, v Value) {
	// Check if this key already exists
	for i, val := range s.keys {
		if s.compare(k, val) == 0 {
			// Overwrite i index of values
			s.values[i] = v
			return
		}
	}

	// Key doesn't exist, append
	s.keys = append(s.keys, k)
	s.values = append(s.values, v)
}

func (s *SymbolTable) Get(k Key) Value {
	// Look for index of k
	for i, val := range s.keys {
		if s.compare(k, val) == 0 {
			// Key found
			return s.values[i]
		}
	}

	// Key not found
	return nil
}

func (s *SymbolTable) Print() {
	// Print all keys and values of this SymbolTable
	for i := range s.keys {
		fmt.Printf("{%v: %v}\n", s.keys[i], s.values[i])
	}
}

package main

import "fmt"

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	var result []T
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	// Contains
	ints := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains(ints, 3)  = %v\n", Contains(ints, 3))
	fmt.Printf("Contains(ints, 99) = %v\n", Contains(ints, 99))

	strs := []string{"go", "is", "fun"}
	fmt.Printf("Contains(strs, \"go\")   = %v\n", Contains(strs, "go"))
	fmt.Printf("Contains(strs, \"rust\") = %v\n", Contains(strs, "rust"))

	// Unique
	dupsInt := []int{1, 2, 2, 3, 3, 3, 4}
	fmt.Printf("\nUnique(%v) = %v\n", dupsInt, Unique(dupsInt))

	dupsStr := []string{"a", "b", "a", "c", "b"}
	fmt.Printf("Unique(%v) = %v\n", dupsStr, Unique(dupsStr))

	// Keys
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("\nKeys(map) = %v\n", Keys(m))
}

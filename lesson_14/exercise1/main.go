package main

import "fmt"

func main()  {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sqr := Map(num, func(n int) int{return n * n})
	evn := Filter(sqr, func(n int) bool {return n %2 == 0})
	sum := Reduce(evn, 0, func(acc, v int) int{return acc + v})

	fmt.Printf("Res: %v\n", sum)
}

func Map[T, R any](slice []T, f func(T) R) []R{
	res := make([]R, len(slice))

	for i, v := range slice {
		res[i] = f(v)
	}

	return res
}

func Filter[T any](slice []T, predicate func(T) bool) []T{
	var res []T
	for _, v := range slice {
		if predicate(v) {
			res = append(res, v)
		}
	}

	return res
}

func Reduce[T, R any](slice []T, initial R, f func(R, T) R) R{
	res := initial
	for _, v := range slice {
		res = f(res, v)
	}

	return res
}
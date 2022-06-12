package main

import "fmt"

func main() {
	fmt.Println(Index[int]([]int{2, 3, 1}, 1))
	fmt.Println(Index[string]([]string{"2", "3", "1"}, "1"))

	fmt.Println(Left[int32]([]int32{2, 3, 1}, 1))
}

func Index[T comparable](s []T, target T) int {
	for i := range s {
		if s[i] == target {
			return i
		}
	}
	return 0
}

type Int interface {
	~int | ~int32 | ~int64
}

func Left[T Int](s []T, target T) int {
	for i := range s {
		if s[i] < target {
			return i
		}
	}
	return 0
}

// 泛型定义及使用
package main

import "fmt"

func main() {
	Print[int]([]int{1, 2, 3})
	fmt.Println(Index[int, string](map[int]string{1: "1", 2: "2"}, 1))
}

func Print[T any](a []T) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func Index[K comparable, V any](m map[K]V, k K) V {
	return m[k]
}

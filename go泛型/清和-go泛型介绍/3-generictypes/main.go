// 泛型类型
// 方法不能包括类型参数
package main

import "fmt"

func main() {
	var s Stack[int]
	s.Push(1)
	fmt.Println(s)
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

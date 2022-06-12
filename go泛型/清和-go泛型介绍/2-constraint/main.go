// 类型约束
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Stringify[I]([]I{1, 2, 3}))
}

type I int

func (i I) String() string {
	return strconv.Itoa(int(i))
}

type Stringer interface {
	String() string
}

func Stringify[T Stringer](s []T) []string {
	var ret = make([]string, 0, len(s))
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

package main

import (
	"math/rand"
	"time"
)

func main() {
	var (
		capacity = 100
		ch       = make(chan uint64, capacity)
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < capacity; i++ {
		ch <- rand.Uint64()
	}

	// 请计算ch内元素值之和
}

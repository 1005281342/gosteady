package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		capacity = 100
		ch       = make(chan uint64, capacity)
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < capacity; i++ {
		//if i == capacity-1 {
		//	ch <- math.MaxUint64	// 一个导致溢出的用例
		//	continue
		//}
		//ch <- uint64(i)
		ch <- rand.Uint64()
	}
	close(ch)

	// 请计算ch内元素值之和
	var (
		wg      = sync.WaitGroup{}
		workers = 3
		goChan  = make(chan struct{}, workers)
		ret     uint64
	)
	for i := 0; i < workers; i++ {
		goChan <- struct{}{}
	}

	for i := 0; i < capacity; i++ {
		<-goChan
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				goChan <- struct{}{}
			}()
			var num, ok = <-ch
			if !ok {
				return
			}
			atomic.AddUint64(&ret, num)
		}()
	}

	wg.Wait()
	close(goChan)
	log.Printf("ret: %d", ret) // 存在的问题：大数溢出
}

// 1. 提升计算速度（并发+同步）
// 1.1 GMP模型
// 2. 考虑溢出，大数相加
// 3. 其他
// 3.1 询问逃逸，为什么在go func() {} 中可以对wg进行操作
// 3.2 代码优化点 channel确定数据生产完毕后，可以进行close

// 查看逃逸情况
// go build -gcflags '-m -l'

package main

import (
	"context"
	"log"
	"time"
)

var (
	ch  chan struct{}
	ch2 chan struct{}
)

// 目的：当chan被关闭时不执行退出，只在超时时退出
// OUT
//2022/05/08 20:04:06 收到了来自ch2的信号
//2022/05/08 20:04:06 收到了来自ch的信号
//2022/05/08 20:04:07 收到了来自ch的信号
//2022/05/08 20:04:08 收到了来自ch的信号
//2022/05/08 20:04:08 收到了来自ch2的信号
//2022/05/08 20:04:10 收到了来自ch2的信号
//2022/05/08 20:04:16 超时退出
func main() {
	ch = make(chan struct{}, 1)
	ch2 = make(chan struct{}, 1)
	go Writer()
	go Writer2()
	Demo()
}

func Writer() {
	for i := 0; i < 3; i++ {
		ch <- struct{}{}
		time.Sleep(time.Second)
	}
	close(ch)
}

func Writer2() {
	for i := 0; i < 3; i++ {
		ch2 <- struct{}{}
		time.Sleep(time.Second * 2)
	}
	close(ch2)
}

func Demo() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case <-ctx.Done(): // case 1
			log.Println("超时退出")
			return
		case _, ok := <-ch: // case 2
			if !ok {
				ch = nil
				continue
			}
			log.Println("收到了来自ch的信号")
		case _, ok := <-ch2: // case 3
			if !ok {
				ch2 = nil
				continue
			}
			log.Println("收到了来自ch2的信号")
		}
	}
}

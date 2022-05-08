package main

import (
	"log"
	"time"
)

func main() {
	//TimerDemo()
	//TimerDemo2()
	//TimerDemo3()
}

func TimerDemo() {
	var timer = time.NewTimer(time.Second)
	for {
		if timer.Stop() {
			return
		}
		log.Println("hi")
	}
}

func TimerDemo2() {
	var timer = time.NewTimer(time.Second)
	for {
		select {
		case _, ok := <-timer.C:
			if !ok {
				return
			}
			log.Println("hi")
		default:
			return
		}
	}
}

func TimerDemo3() {
	for {
		select {
		case _, ok := <-time.After(time.Second):
			if !ok {
				return
			}
			log.Println("hi")
		}
	}
}

func TickerDemo() {
	var ticker = time.NewTicker(time.Second)
	var cnt int
	for range ticker.C { // 从channel中获取数据
		cnt++
		if cnt == 3 {
			ticker.Stop() // 停止计数，timer将不会再定时写入
		}
		log.Println("hi")
	}
}

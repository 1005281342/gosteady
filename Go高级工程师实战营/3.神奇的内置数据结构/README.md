## Chan

```go
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

```



## Timer & Ticker

https://www.cnblogs.com/failymao/p/15073445.html  文中提到了“不再使用的Ticker需要显式地Stop()”，https://www.cnblogs.com/failymao/p/15068712.html 给出了解释：

```go
// Interface to timers implemented in package runtime.
// Must be in sync with ../runtime/time.go:/^type timer
type runtimeTimer struct {
	pp       uintptr
	when     int64
	period   int64  // 当前定时器周期触发间隔（如果是Timer，间隔为0，表示不重复触发）
	f        func(interface{}, uintptr) // NOTE: must not be closure
	arg      interface{}
	seq      uintptr
	nextwhen int64
	status   uint32
}

// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
type Timer struct {
	C <-chan Time
	r runtimeTimer
}


// A Ticker holds a channel that delivers ``ticks'' of a clock
// at intervals.
type Ticker struct {
	C <-chan Time // The channel on which the ticks are delivered.
	r runtimeTimer
}

// 经过一个触发周期时就会将系统时间写入channel
// Timer并不知道什么时候停止计数，所以需要在使用方（读取方）确定不用时（通过Stop方法）主动通知停止计数
```



### 区别

```go
// NewTicker returns a new Ticker containing a channel that will send
// the time on the channel after each tick. The period of the ticks is
// specified by the duration argument. The ticker will adjust the time
// interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will
// panic. Stop the ticker to release associated resources.
func NewTicker(d Duration) *Ticker {
	if d <= 0 {
		panic(errors.New("non-positive interval for NewTicker"))
	}
	// Give the channel a 1-element time buffer.
	// If the client falls behind while reading, we drop ticks
	// on the floor until the client catches up.
	c := make(chan Time, 1)
	t := &Ticker{
		C: c,
		r: runtimeTimer{
			when:   when(d),
			period: int64(d),
			f:      sendTime,
			arg:    c,
		},
	}
	startTimer(&t.r)
	return t
}


// NewTimer creates a new Timer that will send
// the current time on its channel after at least duration d.
func NewTimer(d Duration) *Timer {
	c := make(chan Time, 1)
	t := &Timer{
		C: c,
		r: runtimeTimer{
			when: when(d),
			f:    sendTime,
			arg:  c,
		},
	}
	startTimer(&t.r)
	return t
}

// Ticker会周期触发，而Timer只执行一次
// period字段表示当前定时器周期触发间隔（如果是Timer，间隔为0，表示不重复触发）
// 所以对于Ticker来说需要在确定不再使用定时器的时候主动Stop
```



### 错误的Stop方式导致死锁

```go
package main

import (
	"log"
	"time"
)

func main() {
	go TickerDemo()
	select {}
}

func TickerDemo() {
	var ticker = time.NewTicker(time.Second)
	var cnt int
	for range ticker.C {	// 从channel中获取数据
		cnt++
		if cnt == 3 {
			ticker.Stop() // 停止计数，timer将不会再定时写入
		}
		log.Println("hi")
	}
}

// 对于for range chan语法来说，直到chan被关闭，这个消费者才会结束
// 但是在打印几次之后以及主动停止计数了，将意味着生产者停止了生产
// 思考1：这种1读0写的情况，go语言是怎么检测到并触发panic的？
// 思考2：timer怎么把系统时间写入C（chan Time）的？
// 思考3：ticker中的C有可能被关闭吗？
```


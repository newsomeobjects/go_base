package main

import (
	"context"
	"fmt"
)

// select 可以实现从多个 channel 收发数据，可以使得一个 goroutine 就可以处理多个 channel 的通信。
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()
	for i := 0; i < 2; i++ { //总共处理10次
		select { //每个 case 关键字后面跟的必须是 channel 的发送或者接收操作
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		case z := <-ch3:
			fmt.Printf("receive %d from channel 3\n", z)
			//case ch3 <- 3:
			//	fmt.Printf("send %d to channel 3\n", 3)
			//default:
			//	fmt.Println(";")
		} // 当select没有匹配到case，且没有default： fatal error: all goroutines are asleep - deadlock!
	}

	//在执行 select 语句的时候，如果当下那个时间点没有一个 case 满足条件，就会走 default 分支。
	//
	//至多只能有一个 default 分支。
	//
	//如果没有 default 分支，select 语句就会阻塞，直到某一个 case 满足条件。
	//
	//如果 select 里任何 case 和 default 分支都没有，就会一直阻塞。
	//
	//**如果多个 case 同时满足，select 会随机选一个 case 执行。

	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	//当前main执行完成后，调cancel，然后goroutine停止发消息

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

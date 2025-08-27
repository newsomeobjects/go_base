package main

import (
	"fmt"
	"time"
)

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {

	//var abc chan int
	//ch := make(chan int)，无缓冲，发一个，接一个
	// 创建一个带缓冲的channel，数据满了之后会阻塞
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	//timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭", v)
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-time.After(2 * time.Second): //2s后往这个channel发个消息，接收到消息表示超时
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)              //// 在启动goroutine前调用Add
	c := make(chan int, 2) //最好使用无缓冲通道
	go func() {
		defer wg.Done()
		defer close(c) // 确保通道被关闭
		for i := 1; i < 11; i++ {
			c <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := range c {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}

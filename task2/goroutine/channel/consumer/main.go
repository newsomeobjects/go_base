package main

import (
	"fmt"
	"sync"
)

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Printf("Sent: %d, Buffer usage: %d/%d\n", i, len(ch), 10)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}
